// +build windows

package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/moutend/go-hook/pkg/mouse"
	"github.com/moutend/go-hook/pkg/types"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("error: ")

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// Buffer size is depends on your need. The 100 is placeholder value.
	mouseChan := make(chan types.MouseEvent, 100)

	if err := mouse.Install(nil, mouseChan); err != nil {
		return err
	}

	defer mouse.Uninstall()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	fmt.Println("start capturing mouse input")

	for {
		select {
		case <-time.After(5 * time.Minute):
			fmt.Println("Received timeout signal")
			return nil
		case <-signalChan:
			fmt.Println("Received shutdown signal")
			return nil
		case m := <-mouseChan:
			fmt.Printf("Received %v {X:%v, Y:%v}\n", m.Message, m.X, m.Y)
			continue
		}
	}

	// not reached
	return nil
}
