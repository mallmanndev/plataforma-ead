package models

import "time"

type User struct {
	ID        string    `bson:"_id"`
	Name      string    `bson:"name"`
	Email     string    `bson:"email"`
	Phone     string    `bson:"phone"`
	Type      string    `bson:"type"`
	Approved  bool      `bson:"approved"`
	Password  string    `bson:"password,omitempty"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
