package usecases

import (
	"errors"

	"github.com/matheusvmallmann/plataforma-ead/backend/modules/users/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/users/domain/ports"
	value_objects "github.com/matheusvmallmann/plataforma-ead/backend/modules/users/domain/value-objects"
)

type UpdateUserUseCase struct {
	UsersRepository ports.UsersRepository
}

func NewUpdateUserUseCase(UsersRepository ports.UsersRepository) *UpdateUserUseCase {
	useCase := &UpdateUserUseCase{
		UsersRepository: UsersRepository,
	}

	return useCase
}

type UpdateUserUseCaseInput struct {
	Id    string
	Name  string
	Phone string
}

func (u *UpdateUserUseCase) Execute(Data UpdateUserUseCaseInput) (*entities.User, error) {
	user, err := u.UsersRepository.FindById(Data.Id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("Usuário não encontrado!")
	}

	phone, err := value_objects.NewPhone(Data.Phone)
	if err != nil {
		return nil, err
	}

	user.SetName(Data.Name).SetPhone(phone)

	errUpdate := u.UsersRepository.Update(user)
	if errUpdate != nil {
		return nil, errors.New("Não foi possível atualizar o usuário!!")
	}

	return user, nil
}
