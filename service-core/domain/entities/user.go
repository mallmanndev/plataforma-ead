package entities

import (
	"errors"
	"github.com/google/uuid"
	value_objects "github.com/matheusvmallmann/plataforma-ead/service-core/domain/value-objects"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id        string
	Name      string
	Email     *value_objects.EmailAddress
	Phone     *value_objects.Phone
	Type      *UserType
	Approved  bool
	password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(
	Name string,
	Email *value_objects.EmailAddress,
	Phone *value_objects.Phone,
	Type *UserType,
	Password string,
) (*User, error) {
	user := &User{
		Name:     Name,
		Email:    Email,
		Phone:    Phone,
		Type:     Type,
		Approved: true,
	}

	user.Id = uuid.NewString()
	errPassword := user.ChangePassword(Password)
	errValidate := user.Validate()
	errs := errors.Join(errPassword, errValidate)
	user.CreatedAt = time.Now()

	if errs != nil {
		return nil, errs
	}
	return user, nil
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) ComparePassword(Password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.password), []byte(Password)); err != nil {
		return errors.New("Invalid password!")
	}

	return nil
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("User name invalid!")
	}

	return nil
}

func (u *User) SetName(Name string) *User {
	u.Name = Name
	return u
}

func (u *User) SetPhone(Phone *value_objects.Phone) *User {
	u.Phone = Phone
	return u
}

func (u *User) ChangePassword(Password string) error {
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

func (u *User) SetPassword(Password string) {
	u.password = Password
}
