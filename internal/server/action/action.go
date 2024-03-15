package action

import "time"

type Action struct {
	ID          string
	Description string
	Summary     string
	Type        string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}
