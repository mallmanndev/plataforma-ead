package repositories

import "github.com/matheusvmallmann/plataforma-ead/service-core/domain/entities"

type UsersRepository struct {
}

func NewUsersRepository() *UsersRepository {
	return &UsersRepository{}
}

func (r *UsersRepository) Create(*entities.User) (*entities.User, error) {
	return nil, nil
}

func (r *UsersRepository) Update(*entities.User) (*entities.User, error) {
	return nil, nil
}

func (r *UsersRepository) Delete(*entities.User) error {
	return nil
}

func (r *UsersRepository) FindByEmail(email string) (*entities.User, error) {
	return nil, nil
}

func (r *UsersRepository) FindById(id string) (*entities.User, error) {
	return nil, nil
}
