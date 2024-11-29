package migrations

import (
	"github.com/Amierza/employee-management/migrations/seeds"
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	if err := seeds.ListJobSeeder(db); err != nil {
		return err
	}

	return nil
}
