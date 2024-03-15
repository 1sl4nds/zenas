package comment

import "time"

type Comment struct {
	ID        string
	Body      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
