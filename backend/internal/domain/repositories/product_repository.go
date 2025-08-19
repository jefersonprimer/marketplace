package repositories

import (
	"backend/internal/domain/entities"

	"github.com/google/uuid"
)

type ProductRepository interface {
	Create(product *entities.Product) error
	GetByID(id uuid.UUID) (*entities.Product, error)
	GetBySlug(slug string) (*entities.Product, error)
	Update(product *entities.Product) error
	Delete(id uuid.UUID) error
}
