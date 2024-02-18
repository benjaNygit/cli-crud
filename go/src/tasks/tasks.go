package tasks

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Completed    bool      `json:"completed"`
	Created_at   time.Time `json:"created_at"`
	Completed_at time.Time `json:"completed_at"`
}
