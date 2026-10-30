package repo

import (
	"ecom-project-improved/domain"
	srvProduct "ecom-project-improved/services/product"
	"github.com/jmoiron/sqlx"
)

// ProductRepo interface embeds the contract defined in services/product/port.go.
type ProductRepo interface {
	srvProduct.ProductRepo
}

// productRepo holds the live PostgreSQL database connection handle (*sqlx.DB).
type productRepo struct {
	db *sqlx.DB
}

// NewProductRepo initializes and returns an implementation of the ProductRepo interface.
func NewProductRepo(db *sqlx.DB) ProductRepo {
	repo := &productRepo{
		db: db,
	}
	return repo
}

// Create inserts a new product into PostgreSQL and returns the created product with its generated ID.
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

	// Execute named query
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

func (r *productRepo) List(page, limit int) ([]*domain.Product, error) {
	offset := (page - 1) * limit

	query := `
        SELECT id, title, description, price, img_url
        FROM products
        ORDER BY id
        LIMIT $1 OFFSET $2
    `

	var p []*domain.Product
	err := r.db.Select(&p, query, limit, offset)
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
