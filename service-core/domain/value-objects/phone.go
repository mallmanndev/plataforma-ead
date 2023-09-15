package value_objects

import "errors"

type Phone struct {
	Phone string
}

func NewPhone(phone string) (*Phone, error) {
	phoneStruct := &Phone{
		Phone: phone,
	}
	if isValid := phoneStruct.IsValid(); !isValid {
		return nil, errors.New("Telefone inválido!")
	}

	return phoneStruct, nil
}

func (p *Phone) IsValid() bool {
	if len(p.Phone) < 10 {
		return false
	}

	return true
}
