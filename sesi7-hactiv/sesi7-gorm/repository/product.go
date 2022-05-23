package repository

import (
	"fmt"
	"sesi7-gorm/models"

	"gorm.io/gorm"
)

type ProductRepo interface {
	CreateProduct(*models.Product) error
	GetAllProduct() (*[]models.Product, error)
	GetProductById(id uint) (*models.Product, error)
	DeleteProductById(id uint) error
	UpdateProductById(id uint, name string, brand string) error
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (p *productRepo) CreateProduct(request *models.Product) error {
	err := p.db.Create(request).Error
	return err
}
func (p *productRepo) GetAllProduct() (*[]models.Product, error) {
	var products []models.Product

	err := p.db.Find(&products).Error
	return &products, err
}
func (p *productRepo) GetProductById(id uint) (*models.Product, error) {
	var product models.Product

	err := p.db.First(&product, "id=?", id).Error
	return &product, err
}
func (p *productRepo) DeleteProductById(id uint) error {
	var product models.Product

	err := p.db.Delete(&product, "id=?", id).Error
	return err
}
func (p *productRepo) UpdateProductById(id uint, name string, brand string) error {
	var product models.Product

	err := p.db.Model(&product).Where("id=?", id).Updates(models.Product{Name: name, Brand: brand}).Error
	if err != nil {
		fmt.Println("Error updating data.")
		return err
	}

	fmt.Printf("Success update data with name: %s and brand : %s\n", name, brand)
	return err
}
