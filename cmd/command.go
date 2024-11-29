package cmd

import (
	"log"
	"os"

	"github.com/Amierza/employee-management/migrations"
	"gorm.io/gorm"
)

func Command(db *gorm.DB) {
	migrate := false
	seed := false
	drop := false

	for _, arg := range os.Args[1:] {
		if arg == "--migrate" {
			migrate = true
		}
		if arg == "--seed" {
			seed = true
		}
		if arg == "--drop" {
			drop = true
		}
	}

	if migrate {
		if err := migrations.Migrate(db); err != nil {
			log.Fatal("error migration: %v", err)
		}
		log.Println("migration completed successfully")
	}

	if seed {
		if err := migrations.Seeder(db); err != nil {
			log.Fatal("error migration seeder: %v", err)
		}
		log.Println("seeder completed successfully")
	}

	if drop {
		if err := migrations.DropTables(db); err != nil {
			log.Fatal("error dropping tables: %v", err)
		}
		log.Println("all table dropped successfully")
	}
}
