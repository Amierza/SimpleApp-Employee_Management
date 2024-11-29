package migrations

import (
	"github.com/Amierza/employee-management/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&entity.Employee{},
		&entity.Job{},
	); err != nil {
		return err
	}
	return nil
}
