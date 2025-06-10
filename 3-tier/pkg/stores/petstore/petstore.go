package petstore

import (
	"context"
	"errors"
	"pet-store/pkg/models"
	"pet-store/pkg/stores/database"
)

type PetStore struct {
	cache *database.DBStore
}

func NewPetStore(c *database.DBStore) *PetStore {
	return &PetStore{cache: c}
}

func (p *PetStore) GetPet(ctx context.Context, id int) (*models.Pet, error) {
	pet, err := p.cache.GetPet(id)
	if err != nil {
		return nil, errors.New("pet not found")
	}
	return pet, nil
}

func (p *PetStore) SetPet(ctx context.Context, pet *models.Pet) error {
	return p.cache.SetPet(pet)
}

func (p *PetStore) GetAllPet(ctx context.Context) ([]*models.Pet, error) {
	return p.cache.GetAllPet()
}

func (p *PetStore) Delete(ctx context.Context, id int) error {
	return p.cache.Delete(id)
}
