package usecases

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
)

type UserActivityLogUseCase struct {
	repo repositories.UserActivityLogRepository
}

func NewUserActivityLogUseCase(repo repositories.UserActivityLogRepository) *UserActivityLogUseCase {
	return &UserActivityLogUseCase{repo: repo}
}

func (uc *UserActivityLogUseCase) GetUserActivityLogsByUserId(userId uuid.UUID) ([]*entities.UserActivityLog, error) {
	return uc.repo.GetUserActivityLogsByUserId(userId)
}

func (uc *UserActivityLogUseCase) CreateUserActivityLog(log *entities.UserActivityLog) (*entities.UserActivityLog, error) {
	return uc.repo.Create(log)
}
