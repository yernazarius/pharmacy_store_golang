package services

import (
	"github.com/yernazarius/pharmacy_store_golang/internal/domain/entities"
	"github.com/yernazarius/pharmacy_store_golang/internal/domain/repositories"
)

type ProductService struct {
	Repo *repositories.ProductRepository
}

func (s *ProductService) GetProducts() ([]entities.Product, error) {
	return s.Repo.GetProducts()
}

// Add other CRUD methods similarly...
