package usecases

import (
	"errors"
	"github.com/matheusvmallmann/plataforma-ead/service-core/domain/ports"
	value_objects "github.com/matheusvmallmann/plataforma-ead/service-core/domain/value-objects"

	"github.com/matheusvmallmann/plataforma-ead/service-core/domain/entities"
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

	email, err := value_objects.NewEmailAddress(Email)
	if err != nil {
		return nil, err
	}

	phone, err := value_objects.NewPhone(Phone)
	if err != nil {
		return nil, err
	}

	user, err := entities.NewUser(Name, email, phone, studentUserType, Password)
	if err != nil {
		return nil, err
	}

	findUserByEmail, findUserByEmailErr := u.UsersRepository.FindByEmail(email)
	if findUserByEmailErr != nil {
		return nil, findUserByEmailErr
	}
	if findUserByEmail != nil {
		return nil, errors.New("Email already registered!")
	}

	if err := u.UsersRepository.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}
