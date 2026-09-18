package store

import (
	"backend/models"
	"database/sql"
)

type ScheduleStore interface {
	GetAll() ([]*models.Schedule, error)

	Create(information *models.Schedule) (*models.Schedule, error)
	Update(information *models.Schedule, id int) error
	Delete(id int) error
}

type storeSc struct {
	db *sql.DB
}

func NewScheludeStore(db *sql.DB) ScheduleStore {
	return &storeSc{db: db}
}

func (s *storeSc) GetAll() ([]*models.Schedule, error) {

	q := " SELECT id, day, eventDescription FROM SCHEDULE"
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

func (s *storeSc) Create(information *models.Schedule) (*models.Schedule, error) {
	q := "INSERT INTO SCHEDULE (day, eventDescription) VALUES (?,?)"
	res, err := s.db.Exec(q, information.Day, information.EventDescription)

	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()

	if err != nil {
		return nil, err
	}

	information.Id = int(id)
	return information, nil

}

func (s *storeSc) Update(information *models.Schedule, id int) error {
	q := "UPDATE SCHEDULE SET day = ?, eventDescription = ? WHERE id = ? "
	_, err := s.db.Exec(q, information.Day, information.EventDescription, id)

	if err != nil {
		return err
	}
	return nil

}

func (s *storeSc) Delete(id int) error {
	q := "DELETE FROM SCHEDULE WHERE id = ?"
	_, err := s.db.Exec(q, id)
	if err != nil {
		return err
	}
	return nil
}
