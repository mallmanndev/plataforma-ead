package usecases

import (
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/ports"
)

type CreateCourseUseCase struct {
	courseRepository ports.CourseRepository
}

func NewCreateCourseUseCase(CourseRepository ports.CourseRepository) *CreateCourseUseCase {
	return &CreateCourseUseCase{courseRepository: CourseRepository}
}

type CreateCourseUseCaseDTO struct {
	Name         string
	Description  string
	InstructorId string
}

func (cc *CreateCourseUseCase) Execute(Data CreateCourseUseCaseDTO) (*entities.Course, error) {
	course, err := entities.NewCourse(Data.Name, Data.Description, nil, Data.InstructorId)
	if err != nil {
		return nil, err
	}

	if err := cc.courseRepository.Create(course); err != nil {
		return nil, err
	}

	return course, nil
}
