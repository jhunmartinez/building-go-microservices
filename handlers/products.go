package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PacktPublishing/Building-Microservices-with-Go-Second-Edition/product-api/7_Gorilla/data"
)

type Products struct {
	l *log.Logger
}

// create new function that takes a logger and returns our product handler
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// need our serve HTTP
func (p *Products) ServeHTTP(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProducts()
	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
	rw.Write(d)
}

// got product list want to make a Get Request to ServeHTTP and return the product list.
// we need to look at a package called encoding JSON
// we can convert our product struct into JSON representation
