package handler

import (
	"fmt"
	"pet-store/pkg/models"
	"pet-store/pkg/services"
	"strconv"

	"gofr.dev/pkg/gofr"
)

type PetHandler struct {
	PetService *services.PetService
}

func NewPetHandler(service *services.PetService) *PetHandler {
	return &PetHandler{PetService: service}
}

func (h *PetHandler) GetPets(ctx *gofr.Context) (interface{}, error) {
	pets, err := h.PetService.GetAllPet(ctx.Context)
	if err != nil {
		return nil, err
	}
	if len(pets) == 0 {
		return map[string]string{"message": "No pets found"}, nil
	}
	return pets, nil
}

func (h *PetHandler) GetPetById(ctx *gofr.Context) (interface{}, error) {
	idStr := ctx.PathParam("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, fmt.Errorf("invalid id")
	}

	pet, err := h.PetService.GetPet(ctx.Context, id)
	if err != nil {
		return nil, err
	}
	return pet, nil
}

func (h *PetHandler) CreatePet(ctx *gofr.Context) (interface{}, error) {
	var pet models.Pet
	if err := ctx.Bind(&pet); err != nil {
		return nil, fmt.Errorf("invalid request body")
	}

	err := h.PetService.SetPet(ctx.Context, &pet)
	if err != nil {
		return nil, err
	}
	return pet, nil
}

func (h *PetHandler) DeletePet(ctx *gofr.Context) (interface{}, error) {
	idStr := ctx.PathParam("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, fmt.Errorf("invalid id")
	}

	err = h.PetService.Delete(ctx.Context, id)
	if err != nil {
		return nil, err
	}

	return map[string]string{"message": "Pet deleted successfully"}, nil
}
