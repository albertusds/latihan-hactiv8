package params

type CreateProduct struct {
	Name  string `json:"name"`
	Brand string `json:"brand"`
	Stock int    `json:"stok"`
	Price int    `json:"price"`
}
