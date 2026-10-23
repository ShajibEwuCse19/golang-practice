package repo

import (
	"ecom-project-improved/domain"
	"ecom-project-improved/product"
	"github.com/jmoiron/sqlx"
)

// ProductRepo interface
type ProductRepo interface {
	product.ProductRepo
}

type productRepo struct {
	db *sqlx.DB
}

// This is a constructor whice initialize/create an object of productRepo
func NewProductRepo(db *sqlx.DB) ProductRepo {
	repo := &productRepo{
		db: db,
	}
	return repo
}

func (r *productRepo) Create(p domain.Product) (*domain.Product, error) {
	query := `
	INSERT INTO products (
		title,
		description,
		price,
		img_url
	)
	VALUES (
		:title,
		:description,
		:price,
		:img_url
	)
	RETURNING id
	`

	// Executer named query
	rows, err := r.db.NamedQuery(query, p)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&p.ID)
		if err != nil {
			return nil, err
		}
	}
	return &p, nil
}

func (r *productRepo) List() ([]*domain.Product, error) {
	query := `
	SELECT id, title, description, price, img_url
	FROM products
	`
	var p []*domain.Product
	err := r.db.Select(&p, query)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *productRepo) Get(id uint8) (*domain.Product, error) {
	query := `
	SELECT id, title, description, price, img_url
	FROM products
	WHERE id = $1
	LIMIT 1
	`
	var p domain.Product
	err := r.db.Get(&p, query, id)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *productRepo) Delete(id uint8) error {
	query := `
	DELETE FROM products
	WHERE id = $1
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *productRepo) Update(p domain.Product) (*domain.Product, error) {
	query := `
	UPDATE products
	SET title = :title, description = :description, price = :price, img_url = :img_url
	WHERE id = :id
	RETURNING id
	`
	_, err := r.db.NamedExec(query, p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
