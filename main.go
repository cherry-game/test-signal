package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Program is running, press Ctrl+C or stop debugging to exit...")

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Running...", i)
			time.Sleep(1 * time.Second)
		}
	}()

	<-sigChan
	fmt.Println("Received exit signal, cleaning up resources...")
	doSomeJobs()
	fmt.Println("Program exited gracefully, Bye!")
}

func doSomeJobs() {
	time.Sleep(2 * time.Second)
	fmt.Println("Cleanup completed")
}
