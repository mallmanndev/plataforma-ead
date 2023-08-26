package entities

import (
	"time"

	"github.com/google/uuid"
)

type BaseEntity struct {
	Id        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
