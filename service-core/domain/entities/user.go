package entities

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	BaseEntity
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func NewUser(Name string, Email string, Phone string, Password string) (*User, error) {
	user := &User{
		Name:     Name,
		Email:    Email,
		Phone:    Phone,
		Password: Password,
	}

	user.Id = uuid.New()

	err := user.Validate()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("User name invalid!")
	}
	if u.Email == "" {
		return errors.New("User email invalid!")
	}
	if u.Phone == "" {
		return errors.New("User phone invalid!")
	}
	if u.Password == "" || len(u.Password) < 8 {
		return errors.New("User password invalid!")
	}

	return nil
}
