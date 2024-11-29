package seeds

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/Amierza/employee-management/entity"
	"gorm.io/gorm"
)

func ListEmployeeSeeder(db *gorm.DB) error {
	// membuka file dari format json
	jsonFile, err := os.Open("./migrations/json/employees.json")
	if err != nil {
		return err
	}

	// membaca file json sebagai byte
	jsonData, _ := io.ReadAll(jsonFile)

	// binding dari file json ke struct
	var listEmployee []entity.Employee
	if json.Unmarshal(jsonData, &listEmployee); err != nil {
		return err
	}

	// memastikan tabel employee sudah berada di database jika tidak ada akan membuat tabel employee berdasarkan entity employee
	hasTable := db.Migrator().HasTable(&entity.Employee{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.Employee{}); err != nil {
			return err
		}
	}

	// memasukkan data employee ke database sekaligus melakukan pengecekan terhadap email yang tidak boleh sama
	for _, data := range listEmployee {
		var employee entity.Employee
		err := db.Where(&entity.Employee{Email: data.Email}).First(&employee).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		isData := db.Find(&employee, "email = ?", data.Email).RowsAffected
		if isData == 0 {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
