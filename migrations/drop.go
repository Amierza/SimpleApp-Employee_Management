package migrations

import (
	"gorm.io/gorm"
)

func DropTables(db *gorm.DB) error {
	if err := db.Migrator().DropTable(
		"employees",
		"jobs",
	); err != nil {
		return err
	}

	return nil
}
