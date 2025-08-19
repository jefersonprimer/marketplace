package repositories

import (
	"backend/internal/domain/entities"

	"github.com/google/uuid"
)

type PaymentAttemptRepository interface {
	Create(paymentAttempt *entities.PaymentAttempt) error
	GetByID(id uuid.UUID) (*entities.PaymentAttempt, error)
	GetByPaymentID(paymentID uuid.UUID) ([]*entities.PaymentAttempt, error)
	Update(paymentAttempt *entities.PaymentAttempt) error
	Delete(id uuid.UUID) error
}
