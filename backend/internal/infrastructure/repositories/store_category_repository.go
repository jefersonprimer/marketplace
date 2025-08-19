package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"backend/internal/domain/entities"
)

type StoreCategoryRepositoryImpl struct {
	DB *gorm.DB
}

func (r *StoreCategoryRepositoryImpl) Create(storeCategory *entities.StoreCategory) (*entities.StoreCategory, error) {
	panic("implement me")
}

func (r *StoreCategoryRepositoryImpl) Delete(storeId, categoryId uuid.UUID) error {
	panic("implement me")
}

func (r *StoreCategoryRepositoryImpl) GetCategoriesByStoreId(storeId uuid.UUID) ([]*entities.Category, error) {
	panic("implement me")
}
