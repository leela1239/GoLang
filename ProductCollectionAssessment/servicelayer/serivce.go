package servicelayer

import (
	"product-details/models"
	"product-details/storelayer"
)

type Service struct {
	database *storelayer.Database
}

func NewService(db *storelayer.Database) *Service {
	return &Service{
		database: db,
	}
}

func (s *Service) GetProducts(getquery map[string]string) (*models.Product, error) {
	if getquery["id"] != "" {
		s.database.GetProductsById(getquery["id"])
	} else {
		s.database.GetProductsByName(getquery["name"])
	}

	return nil, nil
}

func (s *Service) PostProducts(products *models.Product) {
	s.database.PostProducts(products)
}

func (s *Service) PutProducts(pid string, products *models.Product) {
	s.database.PutProducts(pid, products)
}

func (s *Service) DeleteProducts(pid string, vid string) {
	if vid == "" {
		s.database.DeleteProducts(pid)
	} else {
		s.database.DeleteVariant(pid, vid)
	}
}
