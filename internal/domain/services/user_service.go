package services

import (
	"github.com/yernazarius/pharmacy_store_golang/internal/domain/entities"
	"github.com/yernazarius/pharmacy_store_golang/internal/domain/repositories"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func (s *UserService) GetUsers() ([]entities.User, error) {
	return s.Repo.GetUsers()
}

// Add other CRUD methods similarly...
