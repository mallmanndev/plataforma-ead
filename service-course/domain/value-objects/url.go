package value_objects

import (
	"net/url"

	errs "github.com/matheusvmallmann/plataforma-ead/service-course/application/errors"
)

type Url struct {
	value string
}

func NewUrl(value string) (*Url, error) {
	url := &Url{
		value: value,
	}
	if err := url.Validate(); err != nil {
		return nil, err
	}
	return url, nil
}

func (u *Url) Validate() error {
	if _, err := url.ParseRequestURI(u.value); err != nil {
		return errs.NewInvalidAttributeError(
			"Url",
			"value",
			"must be a valid url")
	}

	return nil
}

func (u *Url) String() string {
	return u.value
}
