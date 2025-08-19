package repositories

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
)

type WishlistRepository interface {
	FindById(id uuid.UUID) (*entities.Wishlist, error)
	Create(wishlist *entities.Wishlist) (*entities.Wishlist, error)
	GetWishlistsByUserId(userId uuid.UUID) ([]*entities.Wishlist, error)
}
