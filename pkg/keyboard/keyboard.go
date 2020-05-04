// Package keyboard provides low level hook for keyboard input.
package keyboard

import (
	"github.com/moutend/go-hook/pkg/types"
)

// HookHandler is a callback function which processes the incoming low level events.
//
// Note: You don't have to care about this function unless customize the default behavior.
type HookHandler func(c chan<- types.KeyboardEvent) types.HOOKPROC

// Install causes package signal to relay incoming keyboard events to c.
func Install(fn HookHandler, c chan<- types.KeyboardEvent) error {
	return install(fn, c)
}

// Uninstall remove keyboard hook.
func Uninstall() error {
	return uninstall()
}
