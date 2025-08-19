package repositories

import (
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) repositories.ProductRepository {
	return &ProductRepositoryImpl{db: db}
}

func (r *ProductRepositoryImpl) Create(product *entities.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepositoryImpl) GetByID(id uuid.UUID) (*entities.Product, error) {
	var product entities.Product
	err := r.db.Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepositoryImpl) GetBySlug(slug string) (*entities.Product, error) {
	var product entities.Product
	err := r.db.Where("slug = ?", slug).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepositoryImpl) Update(product *entities.Product) error {
	return r.db.Save(product).Error
}

func (r *ProductRepositoryImpl) Delete(id uuid.UUID) error {
	return r.db.Where("id = ?", id).Delete(&entities.Product{}).Error
}
