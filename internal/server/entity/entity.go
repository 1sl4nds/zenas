package entity

import "time"

type Entity struct {
	ID          string
	Name        string
	Description string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}
