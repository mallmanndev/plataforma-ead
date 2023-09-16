package ports

import "github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"

type CourseRepository interface {
	FindById(Id string) (*entities.Course, error)
	Create(Course *entities.Course) error
	Update(Course *entities.Course) error
	Delete(Id string) error
}
