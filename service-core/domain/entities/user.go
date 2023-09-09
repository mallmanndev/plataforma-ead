package entities

import (
	"errors"
	value_objects "github.com/matheusvmallmann/plataforma-ead/service-core/domain/value-objects"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Email     *value_objects.EmailAddress
	Phone     *value_objects.Phone
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(Name string, Email *value_objects.EmailAddress, Phone *value_objects.Phone, Password string) (*User, error) {
	user := &User{
		Id:       uuid.NewString(),
		Name:     Name,
		Email:    Email,
		Phone:    Phone,
		Password: Password,
	}

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
	if u.Password == "" || len(u.Password) < 8 {
		return errors.New("User password invalid!")
	}

	return nil
}
