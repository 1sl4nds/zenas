package address

import "time"

type Address struct {
	ID         string
	Street     string
	Locality   string
	Region     string
	PostalCode string
	Country    string
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	DeletedAt  *time.Time
}
