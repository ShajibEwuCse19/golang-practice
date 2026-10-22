package db

import (
	"ecom-project-improved/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DBConfig struct {
	config config.Config
}

func NewDBConfig(config config.Config) *DBConfig {
	return &DBConfig{
		config: config,
	}
}

func (dbconf *DBConfig) NewConnection() (*sqlx.DB, error) { // *sqlx.DB => connection client
	dbSource := GetConnectionString(dbconf.config)
	dbConnection, err := sqlx.Connect("postgres", dbSource) //there are different driver name for different db. here we are using postgres driver
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return dbConnection, nil // here, dbConnection is client.
}

func GetConnectionString(dbConfig config.Config) string {
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
