package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Employee struct {
	EmployeeID  uuid.UUID `gorm:"type:uuid;primaryKey" json:"employee_id"`
	JobID       uuid.UUID `gorm:"type:uuid;not null" json:"job_id"`
	Job         Job       `gorm:"foreignKey:JobID"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Salary      float64   `json:"salary"`
	TimeStamp
}

func (e *Employee) BeforeCreate(tx *gorm.DB) error {
	e.EmployeeID = uuid.New()
	return nil
}
