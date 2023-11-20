package usecases

import (
	"errors"

	"github.com/matheusvmallmann/plataforma-ead/backend/modules/users/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/users/domain/ports"
	value_objects "github.com/matheusvmallmann/plataforma-ead/backend/modules/users/domain/value-objects"
)

type LoginUseCase struct {
	UsersRepository ports.UsersRepository
}

func NewLoginUseCase(UsersRepository ports.UsersRepository) *LoginUseCase {
	return &LoginUseCase{UsersRepository: UsersRepository}
}

type LoginUseCaseOutput struct {
	User  *entities.User
	Token string
}

func (u *LoginUseCase) Execute(Email string, Password string) (*LoginUseCaseOutput, error) {
	email, err := value_objects.NewEmailAddress(Email)
	if err != nil {
		return nil, err
	}

	user, err := u.UsersRepository.FindByEmail(email)
	if err != nil {
		return nil, errors.New("Error on find user!")
	}

	if user == nil {
		return nil, errors.New("User not found!")
	}

	if err := user.ComparePassword(Password); err != nil {
		return nil, errors.New("Invalid password!")
	}

	return &LoginUseCaseOutput{User: user, Token: "tetete"}, nil
}
