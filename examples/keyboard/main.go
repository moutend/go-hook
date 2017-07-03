// +build windows
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/moutend/go-hook/keyboard"
)

func main() {
	fmt.Println("start capturing keyboard input")

	isInterrupted := false
	ctx, cancel := context.WithCancel(context.Background())

	keyboardChan := make(chan keyboard.KBDLLHOOKSTRUCT, 1)
	keyboard.Notify(ctx, keyboardChan)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	for {
		if isInterrupted {
			cancel()
			break
		}
		select {
		case <-signalChan:
			isInterrupted = true
		case k := <-keyboardChan:
			fmt.Printf("%+v\n", k)
		}
	}
}
