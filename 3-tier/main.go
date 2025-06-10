package main

import (
	"log"
	"pet-store/pkg/config"
	"pet-store/pkg/handler"
	"pet-store/pkg/services"
	"pet-store/pkg/stores/database"
	"pet-store/pkg/stores/petstore"

	"gofr.dev/pkg/gofr"
)

func main() {
	cfg := config.LoadConfig()

	app := gofr.New()

	cache, err := database.NewDBStore(cfg.DBConnectionString)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	store := petstore.NewPetStore(cache)
	service := services.NewPetService(store)
	handler := handler.NewPetHandler(service)

	app.GET("/pets", handler.GetPets)
	app.GET("/pets/{id}", handler.GetPetById)
	app.POST("/pets", handler.CreatePet)
	app.DELETE("/pets/{id}", handler.DeletePet)

	app.Run()
}
