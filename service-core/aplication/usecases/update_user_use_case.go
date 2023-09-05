package usecases

import (
	"errors"

	"github.com/matheusvmallmann/plataforma-ead/service-core/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/service-core/domain/ports"
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

	user.SetName(Data.Name).SetPhone(Data.Phone)

	updatedUser, err := u.UsersRepository.Update(user)
	if err != nil {
		return nil, errors.New("Não foi possível atualizar o usuário!!")
	}

	return updatedUser, nil
}
