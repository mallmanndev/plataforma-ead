package entities

import (
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/mail"
	"time"
)

type User struct {
	BaseEntity
	Name     string
	Email    string
	Phone    string
	Type     *UserType
	Approved bool
	password string
}

func NewUser(Name string, Email string, Phone string, Type *UserType, Password string) (*User, error) {
	user := &User{
		Name:     Name,
		Email:    Email,
		Phone:    Phone,
		Type:     Type,
		Approved: true,
	}

	user.Id = uuid.NewString()
	errPassword := user.SetPassword(Password)
	errValidate := user.ValidateForCreate()
	errs := errors.Join(errPassword, errValidate)
	user.CreatedAt = time.Now()

	if errs != nil {
		return nil, errs
	}
	return user, nil
}

func (u *User) SetPassword(Password string) error {
	if Password == "" || len(Password) < 8 {
		return errors.New("User password invalid!")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.password = string(password)
	return nil
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) ComparePassword(Password string) error {
	return nil
}

func (u *User) ValidateForCreate() error {
	if u.Name == "" {
		return errors.New("User name invalid!")
	}
	if _, err := mail.ParseAddress(u.Email); err != nil {
		return errors.New("User email invalid!")
	}
	if u.Phone == "" || len(u.Phone) < 10 {
		return errors.New("User phone invalid!")
	}

	return nil
}

func (u *User) SetName(Name string) *User {
	u.Name = Name
	return u
}

func (u *User) SetEmail(Email string) *User {
	u.Email = Email
	return u
}

func (u *User) SetPhone(Phone string) *User {
	u.Phone = Phone
	return u
}
