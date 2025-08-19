package repositories

import (
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) repositories.CategoryRepository {
	return &CategoryRepositoryImpl{db: db}
}

func (r *CategoryRepositoryImpl) Create(category *entities.Category) error {
	return r.db.Create(category).Error
}

func (r *CategoryRepositoryImpl) GetByID(id uuid.UUID) (*entities.Category, error) {
	var category entities.Category
	err := r.db.Where("id = ?", id).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepositoryImpl) GetBySlug(slug string) (*entities.Category, error) {
	var category entities.Category
	err := r.db.Where("slug = ?", slug).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepositoryImpl) GetAll() ([]entities.Category, error) {
	var categories []entities.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *CategoryRepositoryImpl) GetByParentID(parentID *uuid.UUID) ([]entities.Category, error) {
	var categories []entities.Category
	query := r.db
	if parentID != nil {
		query = query.Where("parent_id = ?", parentID)
	} else {
		query = query.Where("parent_id IS NULL")
	}
	err := query.Find(&categories).Error
	return categories, err
}

func (r *CategoryRepositoryImpl) Update(category *entities.Category) error {
	return r.db.Save(category).Error
}

func (r *CategoryRepositoryImpl) Delete(id uuid.UUID) error {
	return r.db.Where("id = ?", id).Delete(&entities.Category{}).Error
}
