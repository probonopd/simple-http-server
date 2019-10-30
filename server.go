package main

import (
	"github.com/grandcat/zeroconf"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func serve() {
	http.Handle("/", http.FileServer(http.Dir("/")))
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}

func main() {

	service, err := zeroconf.Register(
		"Go web server", // service instance name
		"_http._tcp",    // service type and protocl
		"local.",        // service domain
		80,              // service port
		nil,             // service metadata
		nil,             // register on all network interfaces
	)

	if err != nil {
		log.Fatal(err)
	}

	defer service.Shutdown()

	go serve()

	log.Println("Serving on port 80")

	// Clean exit.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	select {
	case <-sig:
		// Exit by user
	case <-time.After(time.Second * 120):
		// Exit by timeout
	}

	log.Println("Shutting down")
}
