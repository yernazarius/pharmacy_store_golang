package repositories

import (
	"database/sql"

	"github.com/yernazarius/pharmacy_store_golang/internal/domain/entities"
)

type ProductRepository struct {
	DB *sql.DB
}

func (r *ProductRepository) GetProducts() ([]entities.Product, error) {
	rows, err := r.DB.Query("SELECT id, name, description, price, stock, category_id FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entities.Product
	for rows.Next() {
		var product entities.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CategoryID); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

// Add other CRUD methods similarly...
