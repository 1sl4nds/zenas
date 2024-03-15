package event

import "time"

type Event struct {
	ID          string
	Name        string
	Description string
	StartTime   *time.Time
	EndTime     *time.Time
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}
