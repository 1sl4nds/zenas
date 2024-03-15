package litigation

import "time"

type Litigation struct {
	ID          string
	Name        string
	Description string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}
