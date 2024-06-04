package services

import (
	"context"
	"pharmacy-store/internal/domain/entities"
	"pharmacy-store/internal/domain/repositories"
)

type ProductService struct {
	Repo *repositories.ProductRepository
}

func (s *ProductService) CreateProduct(ctx context.Context, product entities.Product) error {
	return s.Repo.CreateProduct(ctx, product)
}

func (s *ProductService) GetProductByID(ctx context.Context, id int64) (entities.Product, error) {
	return s.Repo.GetProductByID(ctx, id)
}

func (s *ProductService) UpdateProduct(ctx context.Context, product entities.Product) error {
	return s.Repo.UpdateProduct(ctx, product)
}

func (s *ProductService) DeleteProduct(ctx context.Context, id int64) error {
	return s.Repo.DeleteProduct(ctx, id)
}

func (s *ProductService) ListProducts(ctx context.Context, filter string, sort string, limit int, offset int) ([]entities.Product, error) {
	return s.Repo.ListProducts(ctx, filter, sort, limit, offset)
}
