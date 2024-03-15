package completion

import "time"

type Completion struct {
	ID        string
	Content   string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
