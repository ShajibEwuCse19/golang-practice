package repo

import "github.com/jmoiron/sqlx"

type Product struct {
	ID          uint8   `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float32 `json:"price" db:"price"`
	ImgURL      string  `json:"imgURL" db:"img_url"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(id uint8) (*Product, error)
	List() ([]*Product, error)
	Delete(id uint8) error
	Update(p Product) (*Product, error)
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

func (r *productRepo) Create(p Product) (*Product, error) {
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

func (r *productRepo) List() ([]*Product, error) {
	query := `
	SELECT id, title, description, price, img_url
	FROM products
	`
	var p []*Product
	err := r.db.Select(&p, query)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *productRepo) Get(id uint8) (*Product, error) {
	query := `
	SELECT id, title, description, price, img_url
	FROM products
	WHERE id = $1
	LIMIT 1
	`
	var p Product
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

func (r *productRepo) Update(p Product) (*Product, error) {
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
