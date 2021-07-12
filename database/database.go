package database

import (
	"brewery/api/data"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type DBConnection struct {
	l  *log.Logger
	db *sql.DB
}

func NewDBConnection(host string, port int, user string, password string, dbname string, l *log.Logger) (*DBConnection, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	if db.Ping() != nil {
		return nil, err
	}
	return &DBConnection{db: db, l: l}, nil
}

func (db *DBConnection) InsertProduct(name string, desc string, price float32, sku string) error {

	sqlStatement := `
					INSERT INTO products (name, description, price, sku)
					VALUES ($1, $2, $3, $4)`
	_, err := db.db.Exec(sqlStatement, name, desc, price, sku)

	if err != nil {
		db.l.Println(err)
		return err
	}
	return nil
}

func (db *DBConnection) DeleteProductWithID(id int) error {
	sqlStatement := `
					DELETE FROM products
					WHERE id=$1`
	_, err := db.db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}
	return nil
}

func (db *DBConnection) GetAllProducts() error {
	sqlStatement := "SELECT id, name, description, price, sku FROM products"
	results, err := db.db.Query(sqlStatement)
	if err != nil {
		return err
	}
	defer results.Close()

	prod := &data.Product{}
	for results.Next() {
		if err := results.Scan(&prod.ID, &prod.Name, &prod.Description, &prod.Price, &prod.SKU); err != nil {
			db.l.Println(err)
		}
		db.l.Printf("id: %d\nname: %s\ndesc: %s\nprice: %f\nSKU: %s", prod.ID, prod.Name, prod.Description, prod.Price, prod.SKU)
	}
	return nil
}

func (db *DBConnection) UpdateProductWithID(prod data.Product) error {
	sqlStatement := "SELECT sp_update_product($1, $2, $3, $4, $5);"
	_, err := db.db.Exec(sqlStatement, prod.ID, prod.Name, prod.Description, prod.Price, prod.SKU)
	if err != nil {
		db.l.Println(err)
		return err
	}
	return nil
}

/*
func main() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "postgres"
		dbname   = "brewery"
	)
	l := log.New(os.Stdout, "brewery-db ", log.LstdFlags)

	dbConnection, err := NewDBConnection(host, port, user, password, dbname, l)
	if err != nil {
		log.Fatal(err)
	}
	defer dbConnection.db.Close()
	dbConnection.GetAllProducts()
	//dbConnection.InsertProduct("test", "test desc", 1.70, "test")
	//dbConnection.GetAllProducts()
	//dbConnection.DeleteProductWithID(3)
	prod := Product{
		ID:          1,
		Name:        "new new name",
		Description: "new new desc",
		Price:       1,
		SKU:         "new-new-sku",
	}
	dbConnection.UpdateProductWithID(prod)
	dbConnection.GetAllProducts()
}*/
