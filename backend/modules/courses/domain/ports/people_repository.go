package ports

import "github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/entities"

type PeopleRepository interface {
	Upsert(People *entities.People) error
	FindById(Id string) (*entities.People, error)
}
