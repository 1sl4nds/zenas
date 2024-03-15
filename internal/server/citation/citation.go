package citation

import "time"

type Citation struct {
	ID          string
	Description string
	Summary     string
	Type        string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}
