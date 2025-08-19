package repositories

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
)

type StoreCategoryRepository interface {
	Create(storeCategory *entities.StoreCategory) (*entities.StoreCategory, error)
	Delete(storeId, categoryId uuid.UUID) error
	GetCategoriesByStoreId(storeId uuid.UUID) ([]*entities.Category, error)
}
