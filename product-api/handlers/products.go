package handlers

import (
	"log"
	"net/http"

	"github.com/jhunmartinez/building-go-microservices/tree/main/product-api/data"
)

type Products struct {
	l *log.Logger
}

// create new function that takes a logger and returns our product handler
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// need our serve HTTP. Make logic decisions for request
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}
	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// got product list want to make a Get Request to ServeHTTP and return the product list.
// we need to look at a package called encoding JSON
// we can convert our product struct into JSON representation
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts() // GET Request, uses HTTP verb GET
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
