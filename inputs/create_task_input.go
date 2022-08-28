package inputs

import "time"

type CreateTaskInput struct {
	Description string    `json:"description" binding:"required"`
	Deadline    time.Time `json:"deadline" binding:"required"`
	Severity    string    `json:"severity" binding:"required"`
}
