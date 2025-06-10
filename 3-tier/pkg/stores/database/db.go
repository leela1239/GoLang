package database

import (
	"database/sql"
	"errors"
	"pet-store/pkg/models"

	_ "github.com/lib/pq"
)

type DBStore struct {
	db *sql.DB
}

func NewDBStore(connStr string) (*DBStore, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &DBStore{db: db}, nil
}

func (s *DBStore) GetPet(id int) (*models.Pet, error) {
	pet := &models.Pet{}
	query := "SELECT id, name, age FROM pet WHERE id = $1"
	err := s.db.QueryRow(query, id).Scan(&pet.ID, &pet.Name, &pet.Age)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("pet not found")
		}
		return nil, err
	}
	return pet, nil
}

func (s *DBStore) SetPet(pet *models.Pet) error {
	query := `
    INSERT INTO pet ( name, age) 
    VALUES ($1, $2)
    RETURNING id
`
	err := s.db.QueryRow(query, pet.Name, pet.Age).Scan(&pet.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *DBStore) GetAllPet() ([]*models.Pet, error) {
	rows, err := s.db.Query("SELECT id, name, age FROM pet")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pets []*models.Pet
	for rows.Next() {
		pet := &models.Pet{}
		if err := rows.Scan(&pet.ID, &pet.Name, &pet.Age); err != nil {
			return nil, err
		}
		pets = append(pets, pet)
	}

	return pets, nil
}

func (s *DBStore) Delete(id int) error {
	query := "DELETE FROM pet WHERE id = $1"
	_, err := s.db.Exec(query, id)
	return err
}

func (s *DBStore) DeleteAll() error {
	query := "DELETE FROM pet"
	_, err := s.db.Exec(query)
	return err
}
