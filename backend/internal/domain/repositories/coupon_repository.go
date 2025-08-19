
package repositories

import (
    "backend/internal/domain/entities"

    "github.com/google/uuid"
)

type CouponRepository interface {
    Create(coupon *entities.Coupon) error
    GetByID(id uuid.UUID) (*entities.Coupon, error)
    GetByCode(code string) (*entities.Coupon, error)
    Update(coupon *entities.Coupon) error
    Delete(id uuid.UUID) error
}
