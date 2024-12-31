package main

import (
	"fmt"
	server "go-service-template/internal"
	"log"
	"os"
	"os/signal"
	"sync"
)

func main() {
	fmt.Println("Init application")
	defer log.Fatalf("[Info] Application has closed")

	serv, err := server.New()
	if err != nil {
		log.Fatalf("%s", err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		serv.Shutdown()
	}()

	// Wrap each goroutine with a recovery handler
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := serv.App().Listen(serv.Config().Server.Port); err != nil {
			log.Fatalf("%s", err)
		}
	}()

	wg.Wait()
}
