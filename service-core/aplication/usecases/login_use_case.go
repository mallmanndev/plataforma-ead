package usecases

import (
	"errors"

	"github.com/matheusvmallmann/plataforma-ead/service-core/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/service-core/domain/ports"
)

type LoginUseCase struct {
	UsersRepository ports.UsersRepository
}

func (u *LoginUseCase) Execute(Email string, Password string) (*entities.User, error) {
	user, err := u.UsersRepository.FindByEmail(Email)
	if err != nil {
		return nil, errors.New("Error on find user!")
	}

	if user == nil {
		return nil, errors.New("User not found!")
	}

	if err := user.ComparePassword(Password); err != nil {
		return nil, errors.New("Invalid password!")
	}

	return user, nil
}
