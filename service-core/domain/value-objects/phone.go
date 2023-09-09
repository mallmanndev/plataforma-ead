package value_objects

import "errors"

type Phone struct {
	Phone string
}

func NewPhone(PhoneNumber string) (*Phone, error) {
	phone := &Phone{Phone: PhoneNumber}
	if err := phone.Validate(); err != nil {
		return nil, err
	}
	return phone, nil
}

func (p *Phone) Validate() error {
	if len(p.Phone) < 10 {
		return errors.New("Telefone inválido!")
	}

	return nil
}
