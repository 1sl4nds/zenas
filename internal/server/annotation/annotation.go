package annotation

import "time"

type Annotation struct {
	ID            string
	Body          string
	StartPosition string
	EndPosition   string
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	DeletedAt     *time.Time
}
