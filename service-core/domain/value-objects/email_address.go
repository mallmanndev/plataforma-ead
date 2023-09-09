package value_objects

import (
	"errors"
	"regexp"
)

type EmailAddress struct {
	Email string
}

func NewEmailAddress(Email string) (*EmailAddress, error) {
	emailAddess := &EmailAddress{
		Email: Email,
	}

	if err := emailAddess.Validate(); err != nil {
		return nil, err
	}
	return emailAddess, nil
}

func (ea *EmailAddress) Validate() error {
	if len(ea.Email) > 40 {
		return errors.New("Email inválido!")
	}

	if valid, _ := regexp.MatchString(`[^\s]*@[a-z0-9.-]*\.[a-z]{2,6}`, ea.Email); !valid {
		return errors.New("Email inválido!")
	}

	return nil
}
