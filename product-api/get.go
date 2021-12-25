package handlers

import (
	"net/http"

	"github.com/jhunmartinez/building-go-microservices/tree/main/product-api/handlers"
)

func (p *Products) GetProdcts(rw http.ResponseWriter, r *http.Request) }
p.l.Println("Handle GET Products")

// fetch the products from the data store
lp := data.GetProducts()

err := lp.ToJSON(rw)
if err != nil {
	http.Error(rw, "unable to marshal JSON", http.StatusInternalServerError)
}