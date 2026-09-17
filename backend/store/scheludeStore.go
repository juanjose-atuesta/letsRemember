package store

import (
	"backend/models"
	"database/sql"
)


type ScheludeStore interface{
	GetALl() ([]*models.Schedule)
	
	Create(information *models.Schedule) (*models.Schedule, error)
	Update(information *models.Schedule, id int) error
	Delete(id int) error
}



type store struct {
	db *sql.DB
}

func New(db *sql.DB) ScheludeStore {
	return &store{db: db}
}


func (s* store) GetALl() ([]*models.Schedule, error) {

	q := " SELECT id, day, eventDescription FROM SCHELUDE"
	rows, err := s.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

		var information = []*models.Schedule{}
	for rows.Next() {
		b := &models.Schedule{}
		if err := rows.Scan(&b.Id, &b.Day, &b.EventDescription); err != nil {
			return nil, err

		}
		information = append(information, b)
	}
	return information, err


	
}


func (s *store) Create(information *models.Schedule) (*models.Schedule, error){
	q:= ""



}
