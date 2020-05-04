// +build windows

package win32

import (
	"syscall"
	"unsafe"

	"github.com/moutend/go-hook/pkg/types"
)

var (
	modUser32, _ = syscall.LoadDLL("user32.dll")

	procCallNextHookEx, _      = modUser32.FindProc("CallNextHookEx")
	procSetWindowsHookExW, _   = modUser32.FindProc("SetWindowsHookExW")
	procGetMessageW, _         = modUser32.FindProc("GetMessageW")
	procTranslateMessage, _    = modUser32.FindProc("TranslateMessage")
	procDispatchMessageW, _    = modUser32.FindProc("DispatchMessageW")
	procGetModuleHandleW, _    = modUser32.FindProc("GetModuleHandleW")
	procUnhookWindowsHookEx, _ = modUser32.FindProc("UnhookWindowsHookEx")
)

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
