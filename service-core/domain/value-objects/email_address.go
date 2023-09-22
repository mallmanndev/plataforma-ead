package value_objects

import (
	"errors"
	"regexp"
)

type EmailAddress struct {
	Email string
}

func NewEmailAddress(email string) (*EmailAddress, error) {
	emailAddress := &EmailAddress{
		Email: email,
	}
	if isValid := emailAddress.IsValid(); !isValid {
		return nil, errors.New("Email inválido!")
	}

	return emailAddress, nil
}

func (e *EmailAddress) IsValid() bool {
	if len(e.Email) > 40 {
		return false
	}
	valid, _ := regexp.MatchString(`[^\s]*@[a-z0-9.-]*\.[a-z]{2,6}`, e.Email)
	return valid
}
