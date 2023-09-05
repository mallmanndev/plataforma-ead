package ports

import "github.com/matheusvmallmann/plataforma-ead/service-core/domain/entities"

type UsersRepository interface {
	Create(*entities.User) (*entities.User, error)
	Update(*entities.User) (*entities.User, error)
	Delete(*entities.User) error
	FindByEmail(email string) (*entities.User, error)
	FindById(id string) (*entities.User, error)
}
