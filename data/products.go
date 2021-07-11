package data

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
	"io"
	"regexp"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"` //stock-keeping unit
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {

	re := regexp.MustCompile(`[a-z]+-[a-z]+-[1-9]+`)
	matches := re.FindAllString(fl.Field().String(), -1)

	return len(matches) == 1
}

func UpdateProduct(p *Product, id int) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[pos] = p
	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, prod := range productList {
		if prod.ID == id {
			return prod, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

func GetProducts() Products {
	return productList
}

func AddProduct(prod *Product) {
	prod.ID = getNextID()
	productList = append(productList, prod)
}

// temp id gen
func getNextID() int {
	prodList := productList[len(productList)-1]
	return prodList.ID + 1
}

var productList = Products{
	&Product{
		ID:          1,
		Name:        "Luxs",
		Description: "Dit is een amber kleurig bier zacht van afdronk.",
		Price:       1.70,
		SKU:         "luy-luxs-1",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Luxs Classics",
		Description: "TODO",
		Price:       1.70,
		SKU:         "luy-classic-1",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
