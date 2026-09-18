package services

import (
	"backend/models"
	"backend/store"
)

type InformationService struct {
	store store.InformationStore
}

func New(s store.InformationStore) *InformationService {
	return &InformationService{
		store: s,
	}
}

func (s *InformationService) GetAllInformation() ([]*models.Information, error) {
	return s.store.GetAll()
}
func (s *InformationService) GetByTheme(theme string) ([]*models.Information, error) {
	return s.store.GetByTheme(theme)
}

func (s *InformationService) CreateNewInformation(information *models.Information) (*models.Information, error) {
	return s.store.Create(information)

}

func (s *InformationService) UpdateInformation(newInformation *models.Information, id int) error {
	return s.store.Update(newInformation, id)
}

func (s *InformationService) DeleteInformation(id int) error {
	return s.store.Delete(id)
}
