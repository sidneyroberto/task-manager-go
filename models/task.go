package models

import (
	"time"
)

type Task struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	Severity    string    `json:"severity"`
}
