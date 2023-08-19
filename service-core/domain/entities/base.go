package entities

import "time"

type BaseEntity struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt time.Time
}
