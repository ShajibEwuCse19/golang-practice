package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString() string {
	// user -> postgres
	// password -> 123456789
	// host -> localhost
	//port -> 5432 (postgres app setup port)
	// db name -> ecommerce [manually created db name in postgres app]

	return "user=postgres password=123456789 host=localhost port=8080 dbname=ecommerce sslmode=disable"
}

func NewConnection() (*sqlx.DB, error) { // *sqlx.DB => connection client
	dbSource := GetConnectionString()
	dbConnection, err := sqlx.Connect("postgres", dbSource) //there are different driver name for different db. here we are using postgres driver
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return dbConnection, nil // here, dbConnection is client.
}
