package ports

import "github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"

type PeopleRepository interface {
	Create(People *entities.People) error
	Update(People *entities.People) error
	FindById(Id string) (*entities.People, error)
}
