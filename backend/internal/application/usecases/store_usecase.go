package usecases

import (
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"

	"github.com/google/uuid"
)

type StoreUseCase interface {
	CreateStore(store *entities.Store) error
	GetStoreByID(id uuid.UUID) (*entities.Store, error)
	GetStoreBySlug(slug string) (*entities.Store, error)
	GetStoresByOwnerID(ownerID uuid.UUID) ([]entities.Store, error)
	UpdateStore(store *entities.Store) error
	DeleteStore(id uuid.UUID) error
}

type StoreUseCaseImpl struct {
	storeRepo repositories.StoreRepository
}

func NewStoreUseCase(storeRepo repositories.StoreRepository) StoreUseCase {
	return &StoreUseCaseImpl{
		storeRepo: storeRepo,
	}
}

func (uc *StoreUseCaseImpl) CreateStore(store *entities.Store) error {
	return uc.storeRepo.Create(store)
}

func (uc *StoreUseCaseImpl) GetStoreByID(id uuid.UUID) (*entities.Store, error) {
	return uc.storeRepo.GetByID(id)
}

func (uc *StoreUseCaseImpl) GetStoreBySlug(slug string) (*entities.Store, error) {
	return uc.storeRepo.GetBySlug(slug)
}

func (uc *StoreUseCaseImpl) GetStoresByOwnerID(ownerID uuid.UUID) ([]entities.Store, error) {
	return uc.storeRepo.GetByOwnerID(ownerID)
}

func (uc *StoreUseCaseImpl) UpdateStore(store *entities.Store) error {
	return uc.storeRepo.Update(store)
}

func (uc *StoreUseCaseImpl) DeleteStore(id uuid.UUID) error {
	return uc.storeRepo.Delete(id)
}
