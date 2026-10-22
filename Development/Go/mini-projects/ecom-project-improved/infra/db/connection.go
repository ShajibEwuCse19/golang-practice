package db

import (
	"ecom-project-improved/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewConnection(config *config.Config) (*sqlx.DB, error) { // *sqlx.DB => connection client
	dbSource := GetConnectionString(config)
	dbConnection, err := sqlx.Connect("postgres", dbSource) //there are different driver name for different db. here we are using postgres driver
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return dbConnection, nil // here, dbConnection is client.
}

func GetConnectionString(dbConfig *config.Config) string {
	sslMode := "disable"

	if dbConfig.DB.EnableSSlMode {
		sslMode = "require"
	}

	return fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		dbConfig.DB.User,
		dbConfig.DB.Password,
		dbConfig.DB.Host,
		dbConfig.DB.Port,
		dbConfig.DB.Name,
		sslMode,
	)
}
