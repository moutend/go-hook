// +build windows
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/moutend/go-hook/keyboard"
)

func main() {
	fmt.Println("start capturing keyboard input")

	var isInterrupted bool
	var wg sync.WaitGroup

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())
	keyboardChan := make(chan keyboard.KBDLLHOOKSTRUCT, 1)

	go func() {
		wg.Add(1)
		keyboard.Notify(ctx, keyboardChan)
		wg.Done()
	}()
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
	wg.Wait()
	fmt.Println("done")
}
