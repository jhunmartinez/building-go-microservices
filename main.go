package main

import (
	"log"
	"net/http"
	"os"

	"github.com/nicholasjackson/building-microservices-youtube/product-api/handlers"
)

func main() {

	// logger
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	// create reference to the handler
	hh := handlers.NewHello(l)

	// create a new ServeMux now register handler hh to a server
	sm := http.NewServeMux()
	sm.Handle("/", hh)

	http.ListenAndServe(":9090", sm) // server, nil uses default serve mux

}
