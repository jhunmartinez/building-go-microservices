package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/nicholasjackson/building-microservices-youtube/product-api/handlers"
)

func main() {

	// logger
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// create reference to the handler
	// hh := handlers.NewHello(l)
	// gh := handlers.NewGoodbye(l)
	ph := handlers.NewProducts(l)

	// create a new ServeMux now register handler hh to a server
	sm := http.NewServeMux()
	// sm.Handle("/", hh)
	// sm.Handle("/goodbye", gh)
	sm.Handle("/", ph)

	// in depth server, create a new server
	s := &http.Server{
		Addr:         ":9090", // configure the bind addres
		Handler:      sm,      // set the default handler
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		// to run use server.ListenAndServe
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// trap sigterm or interrupt and gracefully shutdown the server
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// Block until a signal is received
	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	// gracefully shutdown the server, waiting max 30 seconds
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

}
