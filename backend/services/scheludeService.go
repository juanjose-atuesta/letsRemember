package services

import (
	"backend/store"
	"baclemd/models"
)

type ScheludeService struct {
	store store.ScheludeStore
}

func New(s *store.ScheludeStore) *ScheludeService {
	return &ScheludeService{
		store: s, 
	}
}

func (s *ScheludeService) GetAllInformation() (*[])
