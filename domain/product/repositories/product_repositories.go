package repositories

import (
	"go-gin-dependency-injection-boilerplate/domain/product/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllProduct() ([]models.Product, error)
	GetProductById(id uint) (*models.Product, error)
	CreateProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
	DeleteProduct(id uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepositroy(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetAllProduct() ([]models.Product, error) {
	var products []models.Product

	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepository) GetProductById(id uint) (*models.Product, error) {
	var product models.Product

	if err := r.db.First(&product, id).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) CreateProduct(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) UpdateProduct(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) DeleteProduct(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}
