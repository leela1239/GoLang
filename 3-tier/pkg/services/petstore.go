package services

import (
	"context"
	"pet-store/pkg/models"
	"pet-store/pkg/stores/petstore"
)

type PetService struct {
	store *petstore.PetStore
}

func NewPetService(store *petstore.PetStore) *PetService {
	return &PetService{store: store}
}

func (s *PetService) GetPet(ctx context.Context, id int) (*models.Pet, error) {
	return s.store.GetPet(ctx, id)
}

func (s *PetService) SetPet(ctx context.Context, pet *models.Pet) error {
	return s.store.SetPet(ctx, pet)
}

func (s *PetService) GetAllPet(ctx context.Context) ([]*models.Pet, error) {
	return s.store.GetAllPet(ctx)
}

func (s *PetService) Delete(ctx context.Context, id int) error {
	return s.store.Delete(ctx, id)
}
