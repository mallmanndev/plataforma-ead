package ports

import (
	"github.com/matheusvmallmann/plataforma-ead/service-core/domain/entities"
	value_objects "github.com/matheusvmallmann/plataforma-ead/service-core/domain/value-objects"
)

type UsersRepository interface {
	Create(*entities.User) error
	Update(*entities.User) error
	Delete(id string) error
	FindByEmail(email *value_objects.EmailAddress) (*entities.User, error)
	FindById(id string) (*entities.User, error)
}
