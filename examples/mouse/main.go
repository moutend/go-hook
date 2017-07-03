// +build windows
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/moutend/go-hook/mouse"
)

func main() {
	fmt.Println("start capturing mouse input")

	isInterrupted := false
	ctx, cancel := context.WithCancel(context.Background())

	mouseChan := make(chan mouse.MSLLHOOKSTRUCT, 1)
	mouse.Notify(ctx, mouseChan)

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
		case k := <-mouseChan:
			fmt.Printf("%+v\n", k)
		}
	}
}
