package repositories

import (
	"database/sql"

	"github.com/yernazarius/pharmacy_store_golang/internal/domain/entities"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) GetUsers() ([]entities.User, error) {
	rows, err := r.DB.Query("SELECT id, name, email, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entities.User
	for rows.Next() {
		var user entities.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// Add other CRUD methods similarly...
