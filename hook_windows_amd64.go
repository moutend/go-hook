// +build windows

// Package hook provides low level hook.
package hook

import (
	"syscall"
	"unsafe"
)

type HOOKPROC uintptr
type HINSTANCE uintptr
type HHOOK uintptr
type HWND uintptr
type LRESULT uintptr

// Point corresponds to Point structure.
// For more information, see the documentation on MSDN.
//
// https://msdn.microsoft.com/en-us/library/windows/desktop/dd162805(v=vs.85).aspx
type POINT struct {
	X int32
	Y int32
}

// MSG corresponds to MSG structure.
// For more information, see the documentation on MSDN.
//
// https://msdn.microsoft.com/en-us/library/windows/desktop/ms644958(v=vs.85).aspx
type MSG struct {
	Hwnd    HWND
	Message uint32
	WParam  uint32
	LParam  uint32
	Time    uint32
	POINT
}

var (
	modUser32, _ = syscall.LoadDLL("user32.dll")

	procCallNextHookEx, _      = modUser32.FindProc("CallNextHookEx")
	procSetWindowsHookExW, _   = modUser32.FindProc("SetWindowsHookExW")
	procGetMessageW, _         = modUser32.FindProc("GetMessageW")
	procGetModuleHandleW, _    = modUser32.FindProc("GetModuleHandleW")
	procUnhookWindowsHookEx, _ = modUser32.FindProc("UnhookWindowsHookEx")
)

func CallNextHookEx(opt, code, wParam, lParam uint64) LRESULT {
	r, _, _ := procCallNextHookEx.Call(
		uintptr(opt),
		uintptr(code),
		uintptr(wParam),
		uintptr(lParam))
	return LRESULT(r)
}

func SetWindowsHookExW(hookId int32, h HOOKPROC, module HINSTANCE, threadId uint32) HHOOK {
	r, _, _ := procSetWindowsHookExW.Call(
		uintptr(hookId),
		uintptr(h),
		uintptr(module),
		uintptr(threadId))
	return HHOOK(r)
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

func UnhookWindowsHookEx(h HHOOK) bool {
	r, _, _ := procUnhookWindowsHookEx.Call(uintptr(h))
	if r == 0 {
		return false
	} else {
		return true
	}
}
