package services

import (
	"backend/models"
	"backend/store"
)

type ScheduleService struct {
	store store.ScheduleStore
}

func NewScheduleService(s store.ScheduleStore) *ScheduleService {
	return &ScheduleService{
		store: s,
	}
}

func (s *ScheduleService) GetAllInformationSchedule() ([]*models.Schedule, error) {
	return s.store.GetAll()
}

func (s *ScheduleService) CreateInformationSchedule(model *models.Schedule) (*models.Schedule, error) {

	return s.store.Create(model)

}

func (s *ScheduleService) UpdateInformationSchedule(update *models.Schedule, id int) error {
	return s.store.Update(update, id)
}

func (s *ScheduleService) DeleteInformationSchedule(id int) error {
	return s.store.Delete(id)
}
