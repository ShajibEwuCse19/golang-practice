module ecom-project-improved

go 1.26.5

require (
	github.com/golang-migrate/migrate/v4 v4.20.1 // db migration
	github.com/jmoiron/sqlx v1.4.0 // execute sql commands
	github.com/lib/pq v1.12.3 // Postgres Database Driver
	github.com/lpernett/godotenv v0.0.0-20230527005122-0de1d4c5ef5e // load .env file into environment variable
	github.com/go-gorp/gorp/v3 v3.1.0 // indirect
	github.com/rakyll/hey v0.1.5 // run and test API performance
	github.com/rubenv/sql-migrate v1.8.1 // indirect
	golang.org/x/net v0.61.0 // indirect
	golang.org/x/text v0.43.0 // indirect
)
