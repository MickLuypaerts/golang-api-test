package data

import (
	"time"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float32
	SKU         string //stock-keeping unit
	CreatedOn   string
	UpdatedOn   string
	DeletedOn   string
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Luxs",
		Description: "Dit is een amber kleurig bier zacht van afdronk.",
		Price:       1.70,
		SKU:         "TODO",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Luxs Classics",
		Description: "TODO",
		Price:       1.70,
		SKU:         "TODO",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
