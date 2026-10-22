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
Here,
migrationSrc,
- We need to provide migration source (directory of migrations sql file)

Exec (Execute Migrations):
--------------------------
db.DB -> db connection and sqlx converts to stdlib *sql.DB (need to verify this)
"postgres" -> database driver name (lib/pq)
migrations -> migrate.FileMigrationSource (migration file source)
migrate.Up -> migration direction (Up, Down, UpAndDown, Reset, N)

return: number of migration applied and error

MigrateDB: It will call when we run go run main.go always
*/