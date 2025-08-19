package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"backend/internal/domain/entities"
)

type ProfileDocumentRepositoryImpl struct {
	DB *gorm.DB
}

func (r *ProfileDocumentRepositoryImpl) FindById(id uuid.UUID) (*entities.ProfileDocument, error) {
	panic("implement me")
}

func (r *ProfileDocumentRepositoryImpl) Create(profileDocument *entities.ProfileDocument) (*entities.ProfileDocument, error) {
	panic("implement me")
}

func (r *ProfileDocumentRepositoryImpl) GetProfileDocumentsByProfileId(profileId uuid.UUID) ([]*entities.ProfileDocument, error) {
	panic("implement me")
}
