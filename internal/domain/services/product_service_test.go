package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yernazarius/pharmacy_store_golang/internal/domain/entities"
	"github.com/yernazarius/pharmacy_store_golang/mocks"
)

func TestGetProducts(t *testing.T) {
	mockRepo := &mocks.ProductRepository{}
	service := &ProductService{Repo: mockRepo}

	mockRepo.On("GetProducts").Return([]entities.Product{
		{ID: 1, Name: "Product 1", Price: 10.0},
		{ID: 2, Name: "Product 2", Price: 20.0},
	}, nil)

	products, err := service.GetProducts()
	assert.NoError(t, err)
	assert.Len(t, products, 2)
	assert.Equal(t, "Product 1", products[0].Name)
}
