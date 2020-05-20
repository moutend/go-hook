package win32

import (
	"syscall"
	"unsafe"

	"github.com/moutend/go-hook/pkg/types"
)

var (
	modUser32, _ = syscall.LoadDLL("user32.dll")

	procGetKeyState, _         = modUser32.FindProc("GetKeyState")
	procToUnicodeEx, _         = modUser32.FindProc("ToUnicodeEx")
	procGetKeyboardLayout, _   = modUser32.FindProc("GetKeyboardLayout")
	procCallNextHookEx, _      = modUser32.FindProc("CallNextHookEx")
	procSetWindowsHookExW, _   = modUser32.FindProc("SetWindowsHookExW")
	procGetMessageW, _         = modUser32.FindProc("GetMessageW")
	procTranslateMessage, _    = modUser32.FindProc("TranslateMessage")
	procDispatchMessageW, _    = modUser32.FindProc("DispatchMessageW")
	procGetModuleHandleW, _    = modUser32.FindProc("GetModuleHandleW")
	procUnhookWindowsHookEx, _ = modUser32.FindProc("UnhookWindowsHookEx")
)

func GetKeyState(nVirtKey int) int {
	r, _, _ := procGetKeyState.Call(uintptr(nVirtKey))

	return int(r)
}

func ToUnicodeEx(wVirtKey, wScanCode int32, dwhkl uintptr) string {
	pwszBuff := make([]uint16, 5)
	lpKeyState := [256]byte{}

	for i, _ := range lpKeyState {
		lpKeyState[i] = byte(GetKeyState(i))
	}

	r, _, _ := procToUnicodeEx.Call(
		uintptr(wVirtKey),
		uintptr(wScanCode),
		uintptr(unsafe.Pointer(&lpKeyState[0])),
		uintptr(unsafe.Pointer(&pwszBuff[0])),
		uintptr(len(pwszBuff)),
		0x4,
		dwhkl,
	)

	if r == 0 {
		return ""
	}

	return syscall.UTF16ToString(pwszBuff)
}

func GetKeyboardLayout(idThread uintptr) uintptr {
	r, _, _ := procGetKeyboardLayout.Call(idThread)

	return r
}

func CallNextHookEx(hhk uintptr, code int32, wParam, lParam uintptr) uintptr {
	r, _, _ := procCallNextHookEx.Call(hhk, uintptr(code), wParam, lParam)

	return r
}

func SetWindowsHookEx(idHook types.Hook, lpfn, hmod uintptr, dwThreadId uint32) uintptr {
	r, _, _ := procSetWindowsHookExW.Call(uintptr(idHook), lpfn, hmod, uintptr(dwThreadId))

	return r
}

func UnhookWindowsHookEx(hhk uintptr) bool {
	r, _, _ := procUnhookWindowsHookEx.Call(hhk)

	if r == 0 {
		return false
	}

	return true
}

func GetMessage(lpMsg **types.MSG, hWnd uintptr, wMsgFilterMin, wMsgFilterMax uint32) int32 {
	r, _, _ := procGetMessageW.Call(
		uintptr(unsafe.Pointer(lpMsg)),
		hWnd,
		uintptr(wMsgFilterMin),
		uintptr(wMsgFilterMin))

	return int32(r)
}

func TranslateMessage(lpMsg **types.MSG) int32 {
	r, _, _ := procTranslateMessage.Call(uintptr(unsafe.Pointer(lpMsg)))

	return int32(r)
}

func DispatchMessage(lpMsg **types.MSG) int32 {
	r, _, _ := procDispatchMessageW.Call(uintptr(unsafe.Pointer(lpMsg)))

	return int32(r)
}

func GetModuleHandle(lpModuleName uintptr) uintptr {
	r, _, _ := procSetWindowsHookExW.Call(lpModuleName)

	return r
}
