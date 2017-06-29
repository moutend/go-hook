// +build windows
package winhook

import (
	"syscall"
	"unsafe"
)

var (
	modUser32, _   = syscall.LoadDLL("user32.dll")
	modKernel32, _ = syscall.LoadDLL("kernel32.dll")

	procCallNextHookEx, _    = modUser32.FindProc("CallNextHookEx")
	procSetWindowsHookExW, _ = modUser32.FindProc("SetWindowsHookExW")
	procGetMessageW, _       = modUser32.FindProc("GetMessageW")
	procGetModuleHandleW, _  = modUser32.FindProc("GetModuleHandleW")
)

func CallNextHookEx(opt, code, wParam, lParam uint32) (lr uintptr) {
	lr, _, _ = procCallNextHookEx.Call(
		uintptr(opt),
		uintptr(code),
		uintptr(wParam),
		uintptr(lParam))
	return
}

func SetWindowsHookExW(hookId int32, fn uintptr, module uintptr, threadId uint32) (hHook uintptr) {
	hHook, _, _ = procSetWindowsHookExW.Call(
		uintptr(hookId),
		uintptr(fn),
		uintptr(module),
		uintptr(threadId))
	return
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
