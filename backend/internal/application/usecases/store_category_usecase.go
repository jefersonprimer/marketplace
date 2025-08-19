package usecases

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
)

type StoreCategoryUseCase struct {
	repo repositories.StoreCategoryRepository
}

func NewStoreCategoryUseCase(repo repositories.StoreCategoryRepository) *StoreCategoryUseCase {
	return &StoreCategoryUseCase{repo: repo}
}

func (uc *StoreCategoryUseCase) GetCategoriesByStoreId(storeId uuid.UUID) ([]*entities.Category, error) {
	return uc.repo.GetCategoriesByStoreId(storeId)
}

func (uc *StoreCategoryUseCase) CreateStoreCategory(storeCategory *entities.StoreCategory) (*entities.StoreCategory, error) {
	return uc.repo.Create(storeCategory)
}

func (uc *StoreCategoryUseCase) DeleteStoreCategory(storeId, categoryId uuid.UUID) error {
	return uc.repo.Delete(storeId, categoryId)
}
