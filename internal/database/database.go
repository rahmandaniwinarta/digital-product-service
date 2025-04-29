package database

import (
	"database/sql"
	"embed"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

// Embed migration files
//
//go:embed migration/*.sql
var dbMigrations embed.FS

var DB *sql.DB

// ConnectDatabase reads DB config from env and opens the connection
func ConnectDatabase() *sql.DB {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	DB = db
	fmt.Println("Database connected successfully!")
	return db
}

// DBMigrate runs SQL migration using embedded .sql files
func DBMigrate(dbParam *sql.DB, direction string) {
	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: dbMigrations,
		Root:       "sql_migrations",
	}

	var migrateDirection migrate.MigrationDirection
	switch direction {
	case "up":
		migrateDirection = migrate.Up
	case "down":
		migrateDirection = migrate.Down
	default:
		fmt.Println("Invalid migration direction. Use 'up' or 'down'.")
		return
	}

	n, err := migrate.Exec(dbParam, "postgres", migrations, migrateDirection)
	if err != nil {
		fmt.Println("Migration failed:", err)

		// Rollback just in case
		n, err := migrate.Exec(dbParam, "postgres", migrations, migrate.Down)
		fmt.Println("Rolled back", n, "migrations.")
		if err != nil {
			panic(err)
		}
		return
	}

	fmt.Println("Migration success, applied", n, "migrations!")
}
