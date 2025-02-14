package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("\n======== STARTING")
	// Create a channel to listen for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	ch := make(chan bool)
	go func() {
		<-quit
		fmt.Print("\nwaiting for a second")
		time.Sleep(time.Millisecond * 250)
		fmt.Print(".")
		time.Sleep(time.Millisecond * 250)
		fmt.Print(".")
		time.Sleep(time.Millisecond * 250)
		fmt.Print(".")
		time.Sleep(time.Millisecond * 250)
		fmt.Println("done!\n=======")
		ch <- true
	}()

	<-ch
}
