package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

type Product struct {
	Id          int
	NameProduct string
	Description string
	Price       float64
	Stock       int
	CodeScan    string
	CategoryId  int
	CreatedAt   time.Time
	UpdateAt    time.Time
}

func InitDB() {
	fmt.Println("Connecting to the database...")
	var err error
	connStr := "host=localhost port=5432 user=postgres password=1234 dbname=ecom_db sslmode=disable"

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database unreachable:", err)
	}

	fmt.Println("Connected to the database successfully!")
}

func GetAllProduct() ([]Product, error) {
	rows, err := DB.Query(`SELECT id, name_product, description, price, stock, code_scan, category_id, created_at, update_at FROM products`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []Product{}
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.Id, &p.NameProduct, &p.Description, &p.Price, &p.Stock, &p.CodeScan, &p.CategoryId, &p.CreatedAt, &p.UpdateAt)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil

}

func main() {

	// cheking connection database
	InitDB()

	fmt.Println("============== GET ALL PRODUCTS ===============")
	productss, err := GetAllProduct()
	if err != nil {
		log.Println("error getting products:", err)
	}

	for _, product := range productss {
		fmt.Printf("ID: %d\n, Name: %s\n, Description: %s\n, Price: %.2f\n, Stock: %d\n, Code Scan: %s\n, Category ID: %d\n, Created At: %s\n, Updated At: %s\n",
			product.Id,
			product.NameProduct,
			product.Description,
			product.Price,
			product.Stock,
			product.CodeScan,
			product.CategoryId,
			product.CreatedAt.Format(time.RFC3339),
			product.UpdateAt.Format(time.RFC3339))
	}

}
