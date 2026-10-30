package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

func MigrateDB(db *sqlx.DB, dir string) error {
	migrationSrc := &migrate.FileMigrationSource{
		Dir: dir,
	}

	_, err := migrate.Exec(db.DB, "postgres", migrationSrc, migrate.Up)
	if err != nil {
		return err
	}

	fmt.Println("Migration completed successfully")
	return nil
}

/*
Migration Notes:
----------------
migrationSrc:
  - FileMigrationSource specifies the directory containing SQL migration files.

migrate.Exec parameters:
  - db.DB        : Passes the underlying standard library *sql.DB handle from *sqlx.DB.
  - "postgres"   : Database driver name (lib/pq).
  - migrationSrc : Migration source pointing to the migrations directory.
  - migrate.Up   : Migration direction. Scans .sql files for "-- +migrate Up" annotations
                   and executes the SQL statements in a database transaction.
*/