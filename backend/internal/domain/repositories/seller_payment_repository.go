package repositories

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
)

type SellerPaymentRepository interface {
	FindById(id uuid.UUID) (*entities.SellerPayment, error)
	Create(sellerPayment *entities.SellerPayment) (*entities.SellerPayment, error)
	GetSellerPaymentsByOrderId(orderId uuid.UUID) ([]*entities.SellerPayment, error)
	GetSellerPaymentsByStoreId(storeId uuid.UUID) ([]*entities.SellerPayment, error)
}
