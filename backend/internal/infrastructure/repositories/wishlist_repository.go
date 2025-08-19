package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"backend/internal/domain/entities"
)

type WishlistRepositoryImpl struct {
	DB *gorm.DB
}

func (r *WishlistRepositoryImpl) FindById(id uuid.UUID) (*entities.Wishlist, error) {
	panic("implement me")
}

func (r *WishlistRepositoryImpl) Create(wishlist *entities.Wishlist) (*entities.Wishlist, error) {
	panic("implement me")
}

func (r *WishlistRepositoryImpl) GetWishlistsByUserId(userId uuid.UUID) ([]*entities.Wishlist, error) {
	panic("implement me")
}
