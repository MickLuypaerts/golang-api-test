// Package classification of Product API
//
// Documentation for Product API
//
//  Schemes: http
//  BasePath: /
//  Version: 1.0.0
//
//  Consumes:
//  - application/json
//
//  Produces:
//  - application/json
//
// swagger:meta
package handlers

import (
	"brewery/api/data"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// A list of products returns in the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// swagger:parameters deleteProduct
type productIDParametersWrapper struct {
	// The id of the product to delete from the database
	// in: path
	// required: true
	ID int `json:"id"`
}

// swagger:response noContent
type productsNoContent struct {
}

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) PUT(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT Products id:", id)
	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	err = data.UpdateProduct(prod, id)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct{}

func (p Products) MiddleWareProductsValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(w, "Error reading product", http.StatusBadRequest)
		}

		// validate the product
		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating product", err)
			http.Error(w, fmt.Sprintf("Error validating product: %s", err), http.StatusBadRequest)
			return
		}

		// add product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
// 	200: productsResponse

// GET returns the products from the data store
func (p *Products) GET(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	prodList := data.GetProducts()
	err := prodList.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json.", http.StatusInternalServerError)
	}
}

func (p *Products) POST(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")

	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	data.AddProduct(prod)
}

// swagger:route DELETE /products/{id} products deleteProduct
// Returns a list of products
// responses:
// 	201: noContent

// DeleteProduct deletes a product from the database
func (p *Products) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle Delete Products id:", id)
	err = data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
}
