package db

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/rlpaul93/order-fulfillment/internal/domain/model"
)

// ProductRepositoryPg implements ProductRepository for PostgreSQL
// Assumes table products(id SERIAL PRIMARY KEY, name TEXT)
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

// PackRepositoryPg implements PackRepository for PostgreSQL
// Assumes table packs(id SERIAL PRIMARY KEY, product_id INT REFERENCES products(id), size INT)
type PackRepositoryPg struct {
	DB *sql.DB
}

func (r *PackRepositoryPg) Create(pack *model.Pack) error {
	return r.DB.QueryRow("INSERT INTO packs(product_id, size) VALUES($1, $2) RETURNING id", pack.ProductID, pack.Size).Scan(&pack.ID)
}

func (r *PackRepositoryPg) GetByID(id uuid.UUID) (*model.Pack, error) {
	p := &model.Pack{}
	row := r.DB.QueryRow("SELECT id, product_id, size FROM packs WHERE id=$1", id)
	if err := row.Scan(&p.ID, &p.ProductID, &p.Size); err != nil {
		return nil, err
	}
	return p, nil
}

func (r *PackRepositoryPg) Update(pack *model.Pack) error {
	_, err := r.DB.Exec("UPDATE packs SET product_id=$1, size=$2 WHERE id=$3", pack.ProductID, pack.Size, pack.ID)
	return err
}

func (r *PackRepositoryPg) Delete(id uuid.UUID) error {
	_, err := r.DB.Exec("DELETE FROM packs WHERE id=$1", id)
	return err
}

func (r *PackRepositoryPg) ListByProduct(productID uuid.UUID) ([]*model.Pack, error) {
	rows, err := r.DB.Query("SELECT id, product_id, size FROM packs WHERE product_id=$1", productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var packs []*model.Pack
	for rows.Next() {
		p := &model.Pack{}
		if err := rows.Scan(&p.ID, &p.ProductID, &p.Size); err != nil {
			return nil, err
		}
		packs = append(packs, p)
	}
	return packs, nil
}
