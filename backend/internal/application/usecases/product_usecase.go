package usecases

import (
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"

	"github.com/google/uuid"
)

type ProductUseCase interface {
	CreateProduct(product *entities.Product) error
	GetProductByID(id uuid.UUID) (*entities.Product, error)
	GetProductBySlug(slug string) (*entities.Product, error)
	GetProductsByCategory(categoryID uuid.UUID) ([]entities.Product, error)
	UpdateProduct(product *entities.Product) error
	DeleteProduct(id uuid.UUID) error
}

type ProductUseCaseImpl struct {
	productRepo repositories.ProductRepository
	categoryRepo repositories.CategoryRepository
	storeRepo repositories.StoreRepository
}

func NewProductUseCase(
	productRepo repositories.ProductRepository,
	categoryRepo repositories.CategoryRepository,
	storeRepo repositories.StoreRepository,
) ProductUseCase {
	return &ProductUseCaseImpl{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
		storeRepo:    storeRepo,
	}
}

func (uc *ProductUseCaseImpl) CreateProduct(product *entities.Product) error {
	// Validar se a categoria existe
	_, err := uc.categoryRepo.GetByID(product.CategoryID)
	if err != nil {
		return err
	}

	// Validar se a loja existe
	_, err = uc.storeRepo.GetByID(product.StoreID)
	if err != nil {
		return err
	}

	return uc.productRepo.Create(product)
}

func (uc *ProductUseCaseImpl) GetProductByID(id uuid.UUID) (*entities.Product, error) {
	return uc.productRepo.GetByID(id)
}

func (uc *ProductUseCaseImpl) GetProductBySlug(slug string) (*entities.Product, error) {
	return uc.productRepo.GetBySlug(slug)
}

func (uc *ProductUseCaseImpl) GetProductsByCategory(categoryID uuid.UUID) ([]entities.Product, error) {
	// TODO: Implementar método no repositório para buscar produtos por categoria
	// Por enquanto, retornamos erro
	return nil, nil
}

func (uc *ProductUseCaseImpl) UpdateProduct(product *entities.Product) error {
	return uc.productRepo.Update(product)
}

func (uc *ProductUseCaseImpl) DeleteProduct(id uuid.UUID) error {
	return uc.productRepo.Delete(id)
}
