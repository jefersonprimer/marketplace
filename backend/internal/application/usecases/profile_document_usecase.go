package usecases

import (
	"github.com/google/uuid"
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
)

type ProfileDocumentUseCase struct {
	repo repositories.ProfileDocumentRepository
}

func NewProfileDocumentUseCase(repo repositories.ProfileDocumentRepository) *ProfileDocumentUseCase {
	return &ProfileDocumentUseCase{repo: repo}
}

func (uc *ProfileDocumentUseCase) GetProfileDocumentsByProfileId(profileId uuid.UUID) ([]*entities.ProfileDocument, error) {
	return uc.repo.GetProfileDocumentsByProfileId(profileId)
}

func (uc *ProfileDocumentUseCase) CreateProfileDocument(document *entities.ProfileDocument) (*entities.ProfileDocument, error) {
	return uc.repo.Create(document)
}
