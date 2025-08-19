package usecases

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
)

type SellerPaymentUseCase struct {
	repo repositories.SellerPaymentRepository
}

func NewSellerPaymentUseCase(repo repositories.SellerPaymentRepository) *SellerPaymentUseCase {
	return &SellerPaymentUseCase{repo: repo}
}

func (uc *SellerPaymentUseCase) GetSellerPaymentsByOrderId(orderId uuid.UUID) ([]*entities.SellerPayment, error) {
	return uc.repo.GetSellerPaymentsByOrderId(orderId)
}

func (uc *SellerPaymentUseCase) GetSellerPaymentsByStoreId(storeId uuid.UUID) ([]*entities.SellerPayment, error) {
	return uc.repo.GetSellerPaymentsByStoreId(storeId)
}

func (uc *SellerPaymentUseCase) CreateSellerPayment(payment *entities.SellerPayment) (*entities.SellerPayment, error) {
	return uc.repo.Create(payment)
}
