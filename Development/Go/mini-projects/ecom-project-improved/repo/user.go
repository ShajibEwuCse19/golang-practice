package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID          uint8  `json:"id" db:"id"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	IsShopOwner bool   `json:"is_shop_owner" db:"is_shop_owner"`
}

type UserRepo interface {
	Create(u User) (*User, error)
	List() ([]*User, error)
	FindUserByEmailAndPassword(email string, password string) (*User, error)
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(user User) (*User, error) {
	query := `
		INSERT INTO users (
		FirstName, 
		LastName, 
		Email, 
		Password, 
		IsShopOwner
		)
		VALUES (
		:first_name, 
		:last_name, 
		:email, 
		:password, 
		:is_shop_owner)
		RETURNING id, created_at
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

func (r *userRepo) List() ([]*User, error) {
	query := `
	SELECT id, first_name, last_name, email, password, is_shop_owner
	FROM users
	`
	var user []*User
	err := r.db.Select(&user, query)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepo) FindUserByEmailAndPassword(email string, password string) (*User, error) {
	query := `
	SELECT id, first_name, last_name, email, password, is_shop_owner
	FROM users
	WHERE email = $1 AND password = $2
	LIMIT 1
	`

	var user User
	err := r.db.Get(&user, query, email, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
