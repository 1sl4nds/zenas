package person

import "time"

type Person struct {
	ID         string
	FamilyName string
	GivenName  string
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	DeletedAt  *time.Time
}
