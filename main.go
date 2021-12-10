package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	// handle func is a convenience method, registers a function a path to a default serve muks
	// default serve mux is an http handler, everything served in Go is an http handler.
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World") // handle func on a path and executes hello world
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Oops", http.StatusBadRequest)
			// rw.WriteHeader(http.StatusBadRequest)
			// rw.Write([]byte("Ooops"))
			return
		}

		// log.Printf("Data %s\n", d)
		// how to write back to user? by using ResponseWriter
		fmt.Fprintf(rw, "Hello %s", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World") // handle func on a path and executes hello world
	})

	http.ListenAndServe(":9090", nil) // server, nil uses default serve mux

	// https://go.dev/src/net/http/server.go?s=61509%3A61556#L2378

	// current program is only logging not really useful
	// how do we read and write to a request? By using ResponseWriter and http.Request

}
