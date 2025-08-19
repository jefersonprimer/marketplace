package usecases

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
)

type StockMovementUseCase struct {
	repo repositories.StockMovementRepository
}

func NewStockMovementUseCase(repo repositories.StockMovementRepository) *StockMovementUseCase {
	return &StockMovementUseCase{repo: repo}
}

func (uc *StockMovementUseCase) GetStockMovementsByProductId(productId uuid.UUID) ([]*entities.StockMovement, error) {
	return uc.repo.GetStockMovementsByProductId(productId)
}

func (uc *StockMovementUseCase) CreateStockMovement(movement *entities.StockMovement) (*entities.StockMovement, error) {
	return uc.repo.Create(movement)
}
