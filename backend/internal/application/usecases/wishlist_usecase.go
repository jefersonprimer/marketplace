package usecases

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
)

type WishlistUseCase struct {
	repo repositories.WishlistRepository
}

func NewWishlistUseCase(repo repositories.WishlistRepository) *WishlistUseCase {
	return &WishlistUseCase{repo: repo}
}

func (uc *WishlistUseCase) GetWishlistsByUserId(userId uuid.UUID) ([]*entities.Wishlist, error) {
	return uc.repo.GetWishlistsByUserId(userId)
}

func (uc *WishlistUseCase) CreateWishlist(wishlist *entities.Wishlist) (*entities.Wishlist, error) {
	return uc.repo.Create(wishlist)
}
