package repositories

import (
	"context"
	"database/sql"
	"pharmacy-store/internal/domain/entities"
)

type ProductRepository struct {
	DB *sql.DB
}

func (r *ProductRepository) CreateProduct(ctx context.Context, product entities.Product) error {
	query := `INSERT INTO products (name, description, price, stock, category_id) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.DB.ExecContext(ctx, query, product.Name, product.Description, product.Price, product.Stock, product.CategoryID)
	return err
}

func (r *ProductRepository) GetProductByID(ctx context.Context, id int64) (entities.Product, error) {
	var product entities.Product
	query := `SELECT id, name, description, price, stock, category_id FROM products WHERE id=$1`
	err := r.DB.QueryRowContext(ctx, query, id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CategoryID)
	return product, err
}

func (r *ProductRepository) UpdateProduct(ctx context.Context, product entities.Product) error {
	query := `UPDATE products SET name=$1, description=$2, price=$3, stock=$4, category_id=$5 WHERE id=$6`
	_, err := r.DB.ExecContext(ctx, query, product.Name, product.Description, product.Price, product.Stock, product.CategoryID, product.ID)
	return err
}

func (r *ProductRepository) DeleteProduct(ctx context.Context, id int64) error {
	query := `DELETE FROM products WHERE id=$1`
	_, err := r.DB.ExecContext(ctx, query, id)
	return err
}

func (r *ProductRepository) ListProducts(ctx context.Context, filter string, sort string, limit int, offset int) ([]entities.Product, error) {
	var products []entities.Product
	query := `SELECT id, name, description, price, stock, category_id FROM products`
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product entities.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CategoryID); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
