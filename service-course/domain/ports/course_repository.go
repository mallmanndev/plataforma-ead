package ports

import "github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"

type GetCourseFilters struct {
	Id      string
	UserId  string
	Visible bool
}

type CourseRepository interface {
	FindById(Id string) (*entities.Course, error)
	Create(Course *entities.Course) error
	Update(Course *entities.Course) error
	Delete(Id string) error
	Get(Filters GetCourseFilters) ([]*entities.Course, error)
	FindBySectionId(SectionId string) (*entities.Course, error)
	FindByItemId(ItemId string) (*entities.Course, error)
}
