package usecases

import (
	errs "github.com/matheusvmallmann/plataforma-ead/service-course/application/errors"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/ports"
)

type CreateSectionUseCase struct {
	coursesRepository ports.CourseRepository
}

func NewCreateSectionUseCase(
	CoursesRepository ports.CourseRepository,
) *CreateSectionUseCase {
	return &CreateSectionUseCase{CoursesRepository}
}

type CreateSectionDTO struct {
	UserId      string
	CourseId    string
	Name        string
	Description string
}

func (cs *CreateSectionUseCase) Execute(Data CreateSectionDTO) (*entities.Course, error) {
	section, err := entities.NewCourseSection(Data.Name, Data.Description, Data.CourseId)
	if err != nil {
		return nil, err
	}

	course, _ := cs.coursesRepository.FindById(Data.CourseId)
	if course == nil {
		return nil, errs.NewNotFoundError("Course")
	}
	if course.InstructorID() != Data.UserId {
		return nil, errs.NewPermissionDeniedError("create section")
	}

	course.AddSection(section)

	if err := cs.coursesRepository.Update(course); err != nil {
		return nil, errs.NewCreateSectionUseCaseError("Could not create section", err)
	}
	return course, nil
}
