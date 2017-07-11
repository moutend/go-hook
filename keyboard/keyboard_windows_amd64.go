// Package keyboard provides low level global hook for keyboard input.
package keyboard

import (
	"context"
	"syscall"
	"unsafe"

	"github.com/moutend/go-hook"
)

// KBDLLHOOKSTRUCT corresponds to KBDLLHOOKSTRUCT structure.
// For more information, see the documentation on MSDN.
//
// https://msdn.microsoft.com/en-us/library/windows/desktop/ms644967(v=vs.85).aspx
type KBDLLHOOKSTRUCT struct {
	VKCode      uint32
	ScanCode    uint32
	Flags       uint32
	Time        uint32
	DWExtraInfo uint32
}

func notify(ctx context.Context, ch chan<- KBDLLHOOKSTRUCT) {
	if ctx == nil {
		panic("hook/keyboard: nil context")
	}
	if ch == nil {
		panic("hook/keyboard: Notify using nil channel")
	}

	const WH_KEYBOARD_LL = 13
	var lResult hook.HHOOK
	hookProcedure := func(code, wParam, lParam uint64) uintptr {
		k := *(*KBDLLHOOKSTRUCT)(unsafe.Pointer(uintptr(lParam)))
		ch <- k
		return uintptr(hook.CallNextHookEx(0, code, wParam, lParam))
	}

	go func() {
		lResult = hook.SetWindowsHookExW(
			WH_KEYBOARD_LL,
			hook.HOOKPROC(syscall.NewCallback(hookProcedure)),
			0,
			0)
		if lResult == 0 {
			panic("failed to set hook procedure")
		}
		var msg *hook.MSG
		hook.GetMessageW(&msg, 0, 0, 0)
	}()

	<-ctx.Done()
	if !hook.UnhookWindowsHookEx(lResult) {
		panic("failed to unhook")
	}
	return
}

// Notify causes package keyboard to relay all keyboard events to ch.
func Notify(ctx context.Context, ch chan<- KBDLLHOOKSTRUCT) {
	notify(ctx, ch)
}
