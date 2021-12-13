package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// add strict interface
type Hello struct {
	l *log.Logger
}

// define function for idiomatic go code
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// add method that satisfies the http handler interface
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	// want to have control over logging at some point for testability
	h.l.Println("Hello World")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello %s", d)

}
