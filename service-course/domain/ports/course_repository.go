package ports

import "github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"

type CourseRepository interface {
	FindById(Id string) (*entities.People, error)
	Create(Course *entities.People) error
	Update(Course *entities.People) error
	Delete(Id string) error
}
