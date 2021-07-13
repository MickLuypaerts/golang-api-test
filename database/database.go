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

func (db *DBConnection) InsertProduct(prod *data.Product) error {

	sqlStatement := `
					INSERT INTO products (name, description, price, sku)
					VALUES ($1, $2, $3, $4)`
	_, err := db.db.Exec(sqlStatement, prod.Name, prod.Description, prod.Price, prod.SKU)

	if err != nil {
		db.l.Println(err)
		return err
	}
	return nil
}

func (db *DBConnection) DeleteProductWithID(id int) error {
	sqlStatement := `
					DELETE FROM 
						products
					WHERE 
						id=$1;`
	results, err := db.db.Exec(sqlStatement, id)
	if err != nil {
		db.l.Println(err)
		return err
	}
	rows, err := results.RowsAffected()
	if err != nil {
		db.l.Println(err)
		return err
	}
	if rows == 0 {
		return data.ErrProductNotFound
	}

	return nil
}

func (db *DBConnection) GetAllProducts() (data.Products, error) {
	sqlStatement := `
					SELECT 
						id, name, description, price, sku 
					FROM 
						products;`
	results, err := db.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer results.Close()
	var prodList data.Products
	for results.Next() {
		prod := &data.Product{}
		if err := results.Scan(&prod.ID, &prod.Name, &prod.Description, &prod.Price, &prod.SKU); err != nil {
			db.l.Println(err)
		}
		prodList = append(prodList, prod)
	}
	return prodList, nil
}

// TODO: Cleanup
func (db *DBConnection) GetProductWithID(id int) (*data.Product, error) {
	db.l.Println("GetProductWithID")
	sqlStatement := `
					SELECT
						id, name, description, price, sku
					FROM
						products
					WHERE
						id=$1;`
	row := db.db.QueryRow(sqlStatement, id)

	prod := &data.Product{}
	err := row.Scan(&prod.ID, &prod.Name, &prod.Description, &prod.Price, &prod.SKU)
	db.l.Printf("%+v\n", prod)
	if err == sql.ErrNoRows {
		return nil, data.ErrProductNotFound
	}
	if err != nil {
		db.l.Println(err)
		return nil, data.ErrProductNotFound
	}
	return prod, nil
}

func (db *DBConnection) UpdateProductWithID(prod *data.Product) error {
	sqlStatement := `
					UPDATE products 
					SET 
						name = $2, description = $3, price = $4, sku = $5, updatedOn = now() 
					WHERE 
						id = $1;`
	results, err := db.db.Exec(sqlStatement, prod.ID, prod.Name, prod.Description, prod.Price, prod.SKU)
	if err != nil {
		db.l.Println(err)
		return err
	}
	rows, err := results.RowsAffected()
	if err != nil {
		db.l.Println(err)
		return err
	}
	if rows == 0 {
		return data.ErrProductNotFound
	}
	return nil
}
