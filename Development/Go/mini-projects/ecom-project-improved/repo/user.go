package repo

import (
	"database/sql"
	"ecom-project-improved/domain"
	"ecom-project-improved/user"
	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	user.UserRepo //Embedding Interface
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(user domain.User) (*domain.User, error) {
	query := `
		INSERT INTO users (
		first_name, 
		last_name, 
		email, 
		password, 
		is_shop_owner
		)
		VALUES (
		:first_name, 
		:last_name, 
		:email, 
		:password, 
		:is_shop_owner)
		RETURNING id
	`

	// Executer named query
	rows, err := r.db.NamedQuery(query, user)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&user.ID)
		if err != nil {
			return nil, err
		}
	}
	return &user, nil
}

func (r *userRepo) List() ([]*domain.User, error) {
	query := `
	SELECT id, first_name, last_name, email, password, is_shop_owner
	FROM users
	`
	var user []*domain.User
	err := r.db.Select(&user, query)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepo) FindUserByEmailAndPassword(email string, password string) (*domain.User, error) {
	query := `
	SELECT id, first_name, last_name, email, password, is_shop_owner
	FROM users
	WHERE email = $1 AND password = $2
	LIMIT 1
	`

	var user domain.User
	err := r.db.Get(&user, query, email, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
