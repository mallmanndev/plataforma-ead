package usecases

import (
	"errors"

	"github.com/matheusvmallmann/plataforma-ead/service-core/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/service-core/domain/ports"
)

type CreateUserUseCase struct {
	UsersRepository ports.UsersRepository
}

func NewCreateUserUseCase(UsersRepository ports.UsersRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		UsersRepository: UsersRepository,
	}
}

func (u *CreateUserUseCase) Execute(Name string, Email string, Phone string, Password string) (*entities.User, error) {
	studentUserType := entities.NewUserType("student", "Student")

	user, err := entities.NewUser(Name, Email, Phone, studentUserType, Password)

	if err != nil {
		return nil, err
	}

	findUserByEmail, findUserByEmailErr := u.UsersRepository.FindByEmail(Email)
	if findUserByEmailErr != nil {
		return nil, findUserByEmailErr
	}
	if findUserByEmail != nil {
		return nil, errors.New("Email already registered!")
	}

	createdUser, err := u.UsersRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
