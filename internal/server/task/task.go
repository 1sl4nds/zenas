package task

import "time"

type Task struct {
	ID          string
	Description string
	Status      string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}
