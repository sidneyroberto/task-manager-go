package outputs

import "task-manager-go/models"

type CreateTaskOutput struct {
	Task models.Task `json:"task"`
}
