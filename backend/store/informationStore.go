package store

import (
	"backend/models"
	"database/sql"
)

type InformationStore interface {
	GetByTheme(theme string) ([]*models.Information, error)
	GetAll() ([]*models.Information, error)
	Create(information *models.Information) (*models.Information, error)
	Update(information *models.Information, id int) error
	Delete(id int) error
}

type store struct {
	db *sql.DB
}

func New(db *sql.DB) InformationStore {
	return &store{db: db}
}

func (s *store) GetAll() ([]*models.Information, error) {
	q := "SELECT theme, description FROM INFORMATION"
	rows, err := s.db.Query(q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var information []*models.Information
	for rows.Next() {
		var b *models.Information
		if err := rows.Scan(&b.Id, &b.Theme, &b.Description); err != nil {
			return nil, err

		}
		information = append(information, b)
	}
	return information, err
}

func (s *store) GetByTheme(theme string) ([]*models.Information, error) {
	q := "SELECT theme, description FROM INFORMATION WHERE theme = ?"
	var information []*models.Information
	rows, err := s.db.Query(q, theme)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var b *models.Information
		if err := rows.Scan(&b.Id, &b.Theme, &b.Description); err != nil {
			return nil, err
		}
		information = append(information, b)

	}
	return information, err

}

func (s *store) Create(information *models.Information) (*models.Information, error) {
	q := "INSERT INTO TABLE INFORMATION (theme, description), VALUES (?,?)"
	resp, err := s.db.Exec(q, information.Theme, information.Description)
	if err != nil {
		return nil, err
	}
	id, err := resp.LastInsertId()

	if err != nil {
		return nil, err
	}

	information.Id = int(id)
	return information, nil
}

func (s *store) Update(information *models.Information, id int) error {
	q := "UPDATE INFORMATION SET theme = ?, description = ? WHERE id = ? "
	_, err := s.db.Exec(q, information.Theme, information.Description)

	if err != nil {
		return err
	}
	return nil

}

func (s *store) Delete(id int) error {
	q := "DELETE FROM INFORMATION WHERE id = ?"
	_, err := s.db.Exec(q, id)
	if err != nil {
		return err
	}
	return nil
}
