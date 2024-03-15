package message

import "time"

type Message struct {
	ID        string
	Body      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
