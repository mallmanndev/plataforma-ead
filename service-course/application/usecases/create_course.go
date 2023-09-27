package usecases

import (
	errs "github.com/matheusvmallmann/plataforma-ead/service-course/application/errors"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/ports"
)

type CreateCourseUseCase struct {
	peopleRepository ports.PeopleRepository
	courseRepository ports.CourseRepository
}

func NewCreateCourseUseCase(
	PeopleRepository ports.PeopleRepository,
	CourseRepository ports.CourseRepository,
) *CreateCourseUseCase {
	return &CreateCourseUseCase{
		peopleRepository: PeopleRepository,
		courseRepository: CourseRepository,
	}
}

type CreateCourseInstructorDTO struct {
	Id   string
	Name string
	Type string
}

type CreateCourseUseCaseDTO struct {
	Name        string
	Description string
	Instructor  CreateCourseInstructorDTO
}

func (cc *CreateCourseUseCase) Execute(Data CreateCourseUseCaseDTO) (*entities.Course, error) {
	people, err := entities.NewPeople(
		Data.Instructor.Id,
		Data.Instructor.Name,
		Data.Instructor.Type,
		nil,
	)
	if err != nil {
		return nil, err
	}

	if err := cc.peopleRepository.Upsert(people); err != nil {
		return nil, errs.NewCreateUserUseCaseError("Could not insert or update people", err)
	}

	course, err := entities.NewCourse(
		Data.Name,
		Data.Description,
		nil,
		Data.Instructor.Id,
	)
	if err != nil {
		return nil, err
	}

	if err := cc.courseRepository.Create(course); err != nil {
		return nil, errs.NewCreateUserUseCaseError("Could not create course", err)
	}

	return course, nil
}
