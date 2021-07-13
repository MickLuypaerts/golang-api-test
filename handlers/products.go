package handlers

import (
	"brewery/api/data"
	"brewery/api/database"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type KeyProduct struct{}

type Products struct {
	l            *log.Logger
	dbConnection *database.DBConnection
}

func NewProducts(handlerLogger *log.Logger, host string, port int, user string, password string, dbname string, dbLogger *log.Logger) *Products {
	dbConn, err := database.NewDBConnection(host, port, user, password, dbname, dbLogger)
	if err != nil {
		dbLogger.Println(err)
	}

	return &Products{l: handlerLogger, dbConnection: dbConn}
}

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
// Returns a list of products from the database
// responses:
// 	200: productsResponse

// ListAll returns all products from the data store
func (p *Products) ListAll(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET all products")

	w.Header().Add("Content-Type", "application/json")
	prodList, err := p.dbConnection.GetAllProducts()
	if err != nil {
		p.l.Println(err)
	}
	err = prodList.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json.", http.StatusInternalServerError)
	}

}

// swagger:route GET /products/{id} products listProduct
// Returns a single product from the database
// responses:
// 	200: productResponse
//	404: description:Product not found

// ListSingle returns a single product from the database
func (p *Products) ListSingle(w http.ResponseWriter, r *http.Request) {
	// TODO: cleanup
	// TODO: swagger doc
	p.l.Println("Handle GET product by id")
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}
	prod, err := p.dbConnection.GetProductWithID(id)
	if err == data.ErrProductNotFound {
		p.l.Println("Product not found in the database")
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
	e := json.NewEncoder(w)
	err = e.Encode(prod)
	if err != nil {
		http.Error(w, "Unable to marshal json.", http.StatusInternalServerError)
	}
}

// swagger:route POST /products products createProduct
// Creates a new product
// responses:
// 	201: description:The product was created successfully.
//	404: description:Product not found

// CreateProduct creates a new product in the database
func (p *Products) POST(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")

	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	p.l.Println(prod)
	p.dbConnection.InsertProduct(prod)
}

// swagger:route DELETE /products/{id} products deleteProduct
// Deletes product with the given id
// responses:
// 	204: description:The product was deleted successfully.
//	404: description:Product not found

// Delete handles DELETE requests and removes items from the database
func (p *Products) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle Delete Products id:", id)
	err = p.dbConnection.DeleteProductWithID(id)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// swagger:route PUT /products{id} products updateProduct
// Updates product with the given id
// responses:
// 	204: description:The product was updated successfully.
//	404: description:Product not found

// UpdateProduct updates a product from the database
func (p *Products) PUT(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT Products id:", id)
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	prod.ID = id

	err = p.dbConnection.UpdateProductWithID(prod)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
