// +build windows
package hook

import (
	"syscall"
	"unsafe"
)

var (
	modUser32, _   = syscall.LoadDLL("user32.dll")
	modKernel32, _ = syscall.LoadDLL("kernel32.dll")

	procCallNextHookEx, _      = modUser32.FindProc("CallNextHookEx")
	procSetWindowsHookExW, _   = modUser32.FindProc("SetWindowsHookExW")
	procGetMessageW, _         = modUser32.FindProc("GetMessageW")
	procGetModuleHandleW, _    = modUser32.FindProc("GetModuleHandleW")
	procUnhookWindowsHookEx, _ = modUser32.FindProc("UnhookWindowsHookEx")
	procPostThreadMessageW, _  = modUser32.FindProc("PostThreadMessageW")
	procPostQuitMessage, _     = modUser32.FindProc("PostQuitMessage")
)

func CallNextHookEx(opt, code, wParam, lParam uint64) (lr uintptr) {
	lr, _, _ = procCallNextHookEx.Call(
		uintptr(opt),
		uintptr(code),
		uintptr(wParam),
		uintptr(lParam))
	return
}

func SetWindowsHookExW(hookId int32, fn uintptr, module uintptr, threadId uint32) HHOOK {
	hHook, _, _ := procSetWindowsHookExW.Call(
		uintptr(hookId),
		uintptr(fn),
		uintptr(module),
		uintptr(threadId))
	return HHOOK(hHook)
}

func GetMessageW(message **MSG, hWindow uintptr, wMsgFilterMin, wMsgFilterMax uint32) bool {
	r, _, _ := procGetMessageW.Call(
		uintptr(unsafe.Pointer(message)),
		hWindow,
		uintptr(wMsgFilterMin),
		uintptr(wMsgFilterMin))
	if r == 0 {
		return false
	} else {
		return true
	}
}

func GetModuleHandleW(name uintptr) (hModule uintptr) {
	hModule, _, _ = procSetWindowsHookExW.Call(
		uintptr(name))
	return
}

func UnhookWindowsHookEx(hHook HHOOK) bool {
	r, _, _ := procUnhookWindowsHookEx.Call(uintptr(hHook))
	if r == 0 {
		return false
	} else {
		return true
	}
}

func PostThreadMessageW(threadId, message, wParam, lParam uint64) bool {
	r, _, _ := procPostThreadMessageW.Call(
		uintptr(threadId),
		uintptr(message),
		uintptr(wParam),
		uintptr(lParam))
	if r == 0 {
		return false
	} else {
		return true
	}
}

func PostQuitMessage(code uint64) {
	procPostQuitMessage.Call(uintptr(code))
}
