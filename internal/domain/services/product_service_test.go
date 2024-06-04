package services

import (
	"context"
	"pharmacy-store/internal/domain/entities"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) CreateProduct(ctx context.Context, product entities.Product) error {
	args := m.Called(ctx, product)
	return args.Error(0)
}

func (m *MockProductRepository) GetProductByID(ctx context.Context, id int64) (entities.Product, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(entities.Product), args.Error(1)
}

func (m *MockProductRepository) UpdateProduct(ctx context.Context, product entities.Product) error {
	args := m.Called(ctx, product)
	return args.Error(0)
}

func (m *MockProductRepository) DeleteProduct(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockProductRepository) ListProducts(ctx context.Context, filter string, sort string, limit int, offset int) ([]entities.Product, error) {
	args := m.Called(ctx, filter, sort, limit, offset)
	return args.Get(0).([]entities.Product), args.Error(1)
}

func TestProductService_CreateProduct(t *testing.T) {
	repo := new(MockProductRepository)
	service := &ProductService{Repo: repo}
	product := entities.Product{Name: "Product1"}

	repo.On("CreateProduct", mock.Anything, product).Return(nil)

	err := service.CreateProduct(context.Background(), product)
	assert.Nil(t, err)
	repo.AssertExpectations(t)
}

// Add more unit tests here
