package repositories

import (
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StoreRepositoryImpl struct {
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) repositories.StoreRepository {
	return &StoreRepositoryImpl{db: db}
}

func (r *StoreRepositoryImpl) Create(store *entities.Store) error {
	return r.db.Create(store).Error
}

func (r *StoreRepositoryImpl) GetByID(id uuid.UUID) (*entities.Store, error) {
	var store entities.Store
	err := r.db.Where("id = ?", id).First(&store).Error
	if err != nil {
		return nil, err
	}
	return &store, nil
}

func (r *StoreRepositoryImpl) GetBySlug(slug string) (*entities.Store, error) {
	var store entities.Store
	err := r.db.Where("slug = ?", slug).First(&store).Error
	if err != nil {
		return nil, err
	}
	return &store, nil
}

func (r *StoreRepositoryImpl) GetByOwnerID(ownerID uuid.UUID) ([]entities.Store, error) {
	var stores []entities.Store
	err := r.db.Where("owner_id = ?", ownerID).Find(&stores).Error
	return stores, err
}

func (r *StoreRepositoryImpl) Update(store *entities.Store) error {
	return r.db.Save(store).Error
}

func (r *StoreRepositoryImpl) Delete(id uuid.UUID) error {
	return r.db.Where("id = ?", id).Delete(&entities.Store{}).Error
}
