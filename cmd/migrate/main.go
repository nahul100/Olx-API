package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/nahul100/Olx-API/internal/config"
	"github.com/nahul100/Olx-API/internal/db"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: migrate <up | down | seed | force>")
	}

	cfg := config.MustLoad()

	m, err := migrate.New(
		"file://internal/migrations",
		cfg.DatabaseURL,
	) //THIS line creates a new instance of the migrate.
	//Migrate struct, which is used to manage database migrations. It takes two arguments: the source of the migration files (in this case, a local file path "file:///migrations") and the database connection URL (cfg.DatabaseURL)
	//that specifies where the migrations should be applied. The migrate.New function initializes the migration process and prepares it for execution.
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}

	switch os.Args[1] {
	case "up":
		if err := m.Up(); err != nil {
			log.Fatal(err)
		}

	case "seed":
		database, err := db.Connect(cfg.DatabaseURL)
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
		defer database.Close()

		seed, err := os.ReadFile("internal/seeds/listings.sql")
		if err != nil {
			log.Fatalf("Failed to read listings seed: %v", err)
		}
		if _, err := database.Exec(string(seed)); err != nil {
			log.Fatalf("Failed to seed listings: %v", err)
		}

	case "down":
		if err := m.Down(); err != nil {
			log.Fatal(err)
		}

	case "force":
		if len(os.Args) < 3 {
			log.Fatal("usage: migrate force <version>")
		}

		var version int
		if _, err := fmt.Sscanf(os.Args[2], "%d", &version); err != nil {
			log.Fatal("invalid version")
		}

		if err := m.Force(version); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("Unknown command: %s (expected up, down, force, or seed)", os.Args[1])
	}

	fmt.Println("Command executed")
}
