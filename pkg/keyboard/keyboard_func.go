// +build !windows

package keyboard

import (
	"fmt"

	"github.com/moutend/go-hook/pkg/types"
)

func install(fn HookHandler, c chan<- types.KeyboardEvent) error {
	return fmt.Errorf("keyboard: not supported")
}

func uninstall() error {
	return fmt.Errorf("keyboard: not supported")
}
