package seeds

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"

	"github.com/Amierza/employee-management/entity"
	"gorm.io/gorm"
)

func ListJobSeeder(db *gorm.DB) error {
	// membuka file format json
	jsonFile, err := os.Open("./migrations/json/jobs.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	// membaca file json sebagai byte
	jsonData, _ := io.ReadAll(jsonFile)

	// binding jsonData ke struct
	var listJobs []entity.Job
	if json.Unmarshal(jsonData, &listJobs); err != nil {
		return err
	}

	// check tabel jobs jika tidak ada maka akan dibuat tabel jobs
	hasTable := db.Migrator().HasTable(&entity.Job{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.Job{}); err != nil {
			return err
		}
	}

	// memasukkan data ke tabel jobs
	for _, data := range listJobs {
		var job entity.Job
		err := db.Where(&entity.Job{JobTitle: data.JobTitle}).First(&job).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Error checking job %s: %v", data.JobTitle, err)
			return err
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := db.Create(&data).Error; err != nil {
				log.Printf("Error inserting job %s: %v", data.JobTitle, err)
				return err
			}
			log.Printf("Inserting new job: %s", data.JobTitle)
		}
	}

	return nil
}
