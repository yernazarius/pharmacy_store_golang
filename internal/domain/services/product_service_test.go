package services

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/yernazarius/pharmacy_store_golang/internal/domain/entities"
	"github.com/yernazarius/pharmacy_store_golang/internal/mocks"
)

func TestGetProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	service := &ProductService{Repo: mockRepo}

	mockRepo.EXPECT().GetProducts().Return([]entities.Product{
		{ID: 1, Name: "Product 1", Price: 10.0},
		{ID: 2, Name: "Product 2", Price: 20.0},
	}, nil)

	products, err := service.GetProducts()
	assert.NoError(t, err)
	assert.Len(t, products, 2)
	assert.Equal(t, "Product 1", products[0].Name)
}
