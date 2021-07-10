package handlers

import (
	"brewery/api/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		p.getProducts(w, r)
	case http.MethodPost:
		p.addProduct(w, r)
	case http.MethodPut:
		// expect the id in th URI
		reg := regexp.MustCompile(`/([0-9]+)`)
		group := reg.FindAllStringSubmatch(r.URL.Path, -1)
		// should only have 1 id
		if len(group) != 1 {
			p.l.Println("Invalid URI more than one id")
			p.l.Println(r.URL.Path)
			p.l.Println(group)
			http.Error(w, "Invalid URI more than one id", http.StatusBadRequest)
			return
		}
		if len(group[0]) != 2 {
			p.l.Println("Invalid URI more than one capture group")
			http.Error(w, "Invalid URI more than one capture group", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(group[0][1])
		if err != nil {
			p.l.Println("Invalid URI unable to convert to int:", group[0][1])
			http.Error(w, "Invalid URI unable to convert to number", http.StatusBadRequest)
			return
		}
		p.updateProduct(w, r, id)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (p *Products) updateProduct(w http.ResponseWriter, r *http.Request, id int) {
	p.l.Println("Handle PUT Products")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to marshal json.", http.StatusBadRequest)
	}
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

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	prodList := data.GetProducts()
	err := prodList.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json.", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to marshal json.", http.StatusBadRequest)
	}
	data.AddProduct(prod)
}
