package models

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name_product"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	CodeScan    string  `json:"code_scan"`
	CategoryID  int     `json:"category_id"`
}

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
