package repositories

import (
	"go-gin-dependency-injection-boilerplate/domain/product/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllProduct([]models.Product) error
	GetProductById(id uint) (*models.Product, error)
	CreateProduct(user *models.Product) error
	UpdateProduct(user *models.Product) error
	DeleteProduct(id uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepositroy(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (p *productRepository) GetAllProduct([]models.Product) error {
	return nil
}
func (p *productRepository) GetProductById(id uint) (*models.Product, error) {
	var product models.Product

	return &product, nil
}
func (p *productRepository) CreateProduct(user *models.Product) error {
	return nil
}
func (p *productRepository) UpdateProduct(user *models.Product) error {
	return nil
}
func (p *productRepository) DeleteProduct(id uint) error {
	return nil
}
