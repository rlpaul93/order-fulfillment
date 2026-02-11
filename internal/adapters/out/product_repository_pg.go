package out

import (
	"database/sql"

	"github.com/rlpaul93/order-fulfillment/internal/domain/model"
)

type ProductRepositoryPg struct {
	DB *sql.DB
}

func (r *ProductRepositoryPg) Create(product *model.Product) error {
	return r.DB.QueryRow("INSERT INTO products(name) VALUES($1) RETURNING id", product.Name).Scan(&product.ID)
}

func (r *ProductRepositoryPg) GetByID(id int64) (*model.Product, error) {
	p := &model.Product{}
	row := r.DB.QueryRow("SELECT id, name FROM products WHERE id=$1", id)
	if err := row.Scan(&p.ID, &p.Name); err != nil {
		return nil, err
	}
	return p, nil
}

func (r *ProductRepositoryPg) Update(product *model.Product) error {
	_, err := r.DB.Exec("UPDATE products SET name=$1 WHERE id=$2", product.Name, product.ID)
	return err
}

func (r *ProductRepositoryPg) Delete(id int64) error {
	_, err := r.DB.Exec("DELETE FROM products WHERE id=$1", id)
	return err
}

func (r *ProductRepositoryPg) List() ([]*model.Product, error) {
	rows, err := r.DB.Query("SELECT id, name FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []*model.Product
	for rows.Next() {
		p := &model.Product{}
		if err := rows.Scan(&p.ID, &p.Name); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
