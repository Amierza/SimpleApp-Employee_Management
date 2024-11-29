package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Job struct {
	JobID    uuid.UUID `gorm:"type:uuid;primaryKey" json:"job_id"`
	JobTitle string    `gorm:"unique;not null" json:"job_title"`
	TimeStamp
}

func (j *Job) BeforeCreate(tx *gorm.DB) error {
	j.JobID = uuid.New()
	return nil
}
