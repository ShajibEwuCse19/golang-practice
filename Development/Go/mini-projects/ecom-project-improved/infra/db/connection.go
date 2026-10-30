package db

import (
	"ecom-project-improved/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// NewConnection establishes a connection pool to PostgreSQL using sqlx.
func NewConnection(config *config.Config) (*sqlx.DB, error) {
	dbSource := GetConnectionString(config)
	dbConnection, err := sqlx.Connect("postgres", dbSource) // Connect using PostgreSQL driver (lib/pq)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return dbConnection, nil
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
