package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"backend/internal/domain/entities"
)

type SellerPaymentRepositoryImpl struct {
	DB *gorm.DB
}

func (r *SellerPaymentRepositoryImpl) FindById(id uuid.UUID) (*entities.SellerPayment, error) {
	panic("implement me")
}

func (r *SellerPaymentRepositoryImpl) Create(sellerPayment *entities.SellerPayment) (*entities.SellerPayment, error) {
	panic("implement me")
}

func (r *SellerPaymentRepositoryImpl) GetSellerPaymentsByOrderId(orderId uuid.UUID) ([]*entities.SellerPayment, error) {
	panic("implement me")
}

func (r *SellerPaymentRepositoryImpl) GetSellerPaymentsByStoreId(storeId uuid.UUID) ([]*entities.SellerPayment, error) {
	panic("implement me")
}
