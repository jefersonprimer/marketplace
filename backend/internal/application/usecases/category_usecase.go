package usecases

import (
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"

	"github.com/google/uuid"
)

type CategoryUseCase interface {
	CreateCategory(category *entities.Category) error
	GetCategoryByID(id uuid.UUID) (*entities.Category, error)
	GetCategoryBySlug(slug string) (*entities.Category, error)
	GetAllCategories() ([]entities.Category, error)
	GetCategoriesByParentID(parentID *uuid.UUID) ([]entities.Category, error)
	UpdateCategory(category *entities.Category) error
	DeleteCategory(id uuid.UUID) error
}

type CategoryUseCaseImpl struct {
	categoryRepo repositories.CategoryRepository
}

func NewCategoryUseCase(categoryRepo repositories.CategoryRepository) CategoryUseCase {
	return &CategoryUseCaseImpl{
		categoryRepo: categoryRepo,
	}
}

func (uc *CategoryUseCaseImpl) CreateCategory(category *entities.Category) error {
	// Validar se a categoria pai existe (se fornecida)
	if category.ParentID != uuid.Nil {
		_, err := uc.categoryRepo.GetByID(category.ParentID)
		if err != nil {
			return err
		}
	}

	return uc.categoryRepo.Create(category)
}

func (uc *CategoryUseCaseImpl) GetCategoryByID(id uuid.UUID) (*entities.Category, error) {
	return uc.categoryRepo.GetByID(id)
}

func (uc *CategoryUseCaseImpl) GetCategoryBySlug(slug string) (*entities.Category, error) {
	return uc.categoryRepo.GetBySlug(slug)
}

func (uc *CategoryUseCaseImpl) GetAllCategories() ([]entities.Category, error) {
	return uc.categoryRepo.GetAll()
}

func (uc *CategoryUseCaseImpl) GetCategoriesByParentID(parentID *uuid.UUID) ([]entities.Category, error) {
	return uc.categoryRepo.GetByParentID(parentID)
}

func (uc *CategoryUseCaseImpl) UpdateCategory(category *entities.Category) error {
	return uc.categoryRepo.Update(category)
}

func (uc *CategoryUseCaseImpl) DeleteCategory(id uuid.UUID) error {
	return uc.categoryRepo.Delete(id)
}
