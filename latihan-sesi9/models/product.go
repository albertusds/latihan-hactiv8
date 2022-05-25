package models

import "fmt"

type Product struct {
	Name  string `json:"name"`
	Brand string `json:"brand"`
	Stock int    `json:"stok"`
	Price int    `json:"price"`
}

var Products []Product

func GetProducts() *[]Product {
	return &Products
}

func GetProductsByBrand(brand string) (Product, error) {
	var p Product
	isExist := false

	for _, product := range Products {
		if product.Brand == brand {
			p = product
			isExist = true
			break
		}
	}

	if !isExist {
		return Product{}, fmt.Errorf("no product with brand %s", brand)
	}

	return p, nil
}
