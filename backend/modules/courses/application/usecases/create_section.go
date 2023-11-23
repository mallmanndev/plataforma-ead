package usecases

import (
	errs "github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/errors"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/ports"
	value_objects "github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/value-objects"
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
	UserId        string
	CourseId      string
	Name          string
	Description   string
	AvaliationUrl string
}

func (cs *CreateSectionUseCase) Execute(Data CreateSectionDTO) (*entities.Course, error) {
	avaliation, err := cs.getAvaliationUrl(Data.AvaliationUrl)
	if err != nil {
		return nil, err
	}

	section, err := entities.NewCourseSection(Data.Name, Data.Description, Data.CourseId, avaliation)
	if err != nil {
		return nil, err
	}

	course, _ := cs.coursesRepository.FindById(Data.CourseId)
	if course == nil {
		return nil, errs.NewNotFoundError("Course")
	}
	if course.UserId() != Data.UserId {
		return nil, errs.NewPermissionDeniedError("create section")
	}

	course.AddSection(section)

	if err := cs.coursesRepository.Update(course); err != nil {
		return nil, errs.NewCreateSectionUseCaseError("Could not create section", err)
	}
	return course, nil
}

func (cs *CreateSectionUseCase) getAvaliationUrl(url string) (*value_objects.Url, error) {
	if url == "" {
		return nil, nil
	}

	discord, err := value_objects.NewUrl(url)
	if err != nil {
		return nil, err
	}

	return discord, nil
}
