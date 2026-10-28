package main

import (
	"fmt"
	"log"
	"os"

	"github.com/anoopraju31/olx-api/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) < 2 {
		log.Fatal("USEAGE: migrate <up | down>")
	}

	cfg := config.MustLoad()

	m, migErr := migrate.New(
		"file://migrations",
		cfg.DatabaseUrl,
	)
	if migErr != nil {
		log.Fatalf("migration.new: %v", migErr)
	}

	switch os.Args[1] {
	case "up":
		if err := m.Up(); err != nil {
			log.Fatal((err))
		}
		log.Println("UP")
	case "down":
		if err := m.Down(); err != nil {
			log.Fatal((err))
		}
		log.Println("DOWN")
	default:
		log.Fatalf("UNKNOW COMMAND: %s", os.Args[1])
	}
}
