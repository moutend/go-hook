// +build windows

package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
	"unsafe"

	"github.com/micmonay/keybd_event"
	"github.com/moutend/go-hook/pkg/keyboard"
	"github.com/moutend/go-hook/pkg/types"
	"github.com/moutend/go-hook/pkg/win32"
)

var (
	kbA, kbB keybd_event.KeyBonding
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("error: ")

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() (err error) {
	// Emulates pressing and releasing the 'A' key.
	kbA, err = keybd_event.NewKeyBonding()

	if err != nil {
		return err
	}

	kbA.SetKeys(keybd_event.VK_A)

	// Emulates pressing and releasing the 'B' key.
	kbB, err = keybd_event.NewKeyBonding()

	if err != nil {
		return err
	}

	kbB.SetKeys(keybd_event.VK_B)

	// Buffer size is depends on your need. The 100 is placeholder value.
	keyboardChan := make(chan types.KeyboardEvent, 100)

	if err := keyboard.Install(handler, keyboardChan); err != nil {
		return err
	}

	defer keyboard.Uninstall()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	fmt.Println("start capturing keyboard input")

	for {
		select {
		case <-time.After(5 * time.Minute):
			fmt.Println("Received timeout signal")
			return nil
		case <-signalChan:
			fmt.Println("Received shutdown signal")
			return nil
		case k := <-keyboardChan:
			fmt.Printf("Received %V %v\n", k.Message, k.VKCode)
			continue
		}
	}

	// not reached
	return nil
}

func handler(c chan<- types.KeyboardEvent) types.HOOKPROC {
	counter := 0

	return func(code int32, wParam, lParam uintptr) uintptr {
		if lParam == 0 {
			goto NEXT
		}

		c <- types.KeyboardEvent{
			Message:         types.Message(wParam),
			KBDLLHOOKSTRUCT: *(*types.KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam)),
		}

		switch (*types.KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam)).VKCode {
		case types.VK_A:
			if counter == 1 {
				counter = 0
				goto NEXT
			}
			switch types.Message(wParam) {
			case types.WM_KEYDOWN:
				go kbB.Press()
			case types.WM_KEYUP:
				go kbB.Release()
			}

			counter = 1

			return 1
		case types.VK_B:
			if counter == 1 {
				counter = 0
				goto NEXT
			}
			switch types.Message(wParam) {
			case types.WM_KEYDOWN:
				go kbA.Press()
			case types.WM_KEYUP:
				go kbA.Release()
			}

			counter = 1

			return 1
		default:
		}

	NEXT:

		return win32.CallNextHookEx(0, code, wParam, lParam)
	}
}
