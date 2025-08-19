package repositories

import (
	"backend/internal/domain/entities"

	"github.com/google/uuid"
)

type PaymentRepository interface {
	Create(payment *entities.Payment) error
	GetByID(id uuid.UUID) (*entities.Payment, error)
	GetByOrderGroupID(orderGroupID uuid.UUID) (*entities.Payment, error)
	Update(payment *entities.Payment) error
	Delete(id uuid.UUID) error
}
