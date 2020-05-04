// Package mouse provides low level hook for mouse input.
package mouse

import (
	"github.com/moutend/go-hook/pkg/types"
)

// HookHandler is a callback function which processes the incoming low level events.
//
// Note: You don't have to care about this function unless customize the default behavior.
type HookHandler func(c chan<- types.MouseEvent) types.HOOKPROC

// Install causes package signal to relay incoming mouse events to c.
func Install(fn HookHandler, c chan<- types.MouseEvent) error {
	return install(fn, c)
}

// Uninstall remove mouse hook.
func Uninstall() error {
	return uninstall()
}
