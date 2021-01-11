package main

import (
	"fmt"
	"godemo/src/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	http.Start()
	ending()
}

func ending() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	case <-c:
		fmt.Printf("stop signal caught, stopping... pid=%d\n", os.Getpid())
	}
	http.Shutdown()
	fmt.Println("godemo stopped successfully")
}
