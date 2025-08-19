package usecases

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
)

type ReturnUseCase struct {
	repo repositories.ReturnRepository
}

func NewReturnUseCase(repo repositories.ReturnRepository) *ReturnUseCase {
	return &ReturnUseCase{repo: repo}
}

func (uc *ReturnUseCase) GetReturnsByOrderId(orderId uuid.UUID) ([]*entities.Return, error) {
	return uc.repo.GetReturnsByOrderId(orderId)
}

func (uc *ReturnUseCase) CreateReturn(returnEntity *entities.Return) (*entities.Return, error) {
	return uc.repo.Create(returnEntity)
}
