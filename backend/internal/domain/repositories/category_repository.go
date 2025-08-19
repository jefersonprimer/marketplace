
package repositories

import (
    "backend/internal/domain/entities"

    "github.com/google/uuid"
)

type CategoryRepository interface {
    Create(category *entities.Category) error
    GetByID(id uuid.UUID) (*entities.Category, error)
    GetBySlug(slug string) (*entities.Category, error)
    GetAll() ([]entities.Category, error)
    GetByParentID(parentID *uuid.UUID) ([]entities.Category, error)
    Update(category *entities.Category) error
    Delete(id uuid.UUID) error
}
