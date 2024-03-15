package document

import "time"

type Document struct {
	ID          string
	Title       string
	Description string
	URL         string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}
