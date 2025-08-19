
package repositories

import (
    "backend/internal/domain/entities"

    "github.com/google/uuid"
)

type CartItemRepository interface {
    Create(cartItem *entities.CartItem) error
    GetByID(id uuid.UUID) (*entities.CartItem, error)
    GetByUserID(userID uuid.UUID) ([]*entities.CartItem, error)
    Update(cartItem *entities.CartItem) error
    Delete(id uuid.UUID) error
}
