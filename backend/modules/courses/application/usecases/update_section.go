package usecases

import (
	errs "github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/errors"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/ports"
	value_objects "github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/value-objects"
)

type UpdateSectionUseCase struct {
	coursesRepository ports.CourseRepository
}

func NewUpdateSectionUseCase(
	CoursesRepository ports.CourseRepository,
) *UpdateSectionUseCase {
	return &UpdateSectionUseCase{CoursesRepository}
}

type UpdateSectionDTO struct {
	UserId        string
	SectionId     string
	Name          string
	Description   string
	AvaliationUrl string
}

func (cs *UpdateSectionUseCase) Execute(Data UpdateSectionDTO) (*entities.Course, error) {
	course, _ := cs.coursesRepository.FindBySectionId(Data.SectionId)
	if course == nil {
		return nil, errs.NewUpdateSectionUseCaseError("Course not found", nil)
	}
	if course.UserId() != Data.UserId {
		return nil, errs.NewPermissionDeniedError("update section")
	}

	section := course.FindSection(Data.SectionId)

	avaliation, err := cs.getAvaliationUrl(Data.AvaliationUrl)
	if err != nil {
		return nil, err
	}

	if err := section.Update(Data.Name, Data.Description, avaliation); err != nil {
		return nil, err
	}

	if err := cs.coursesRepository.Update(course); err != nil {
		return nil, errs.NewUpdateSectionUseCaseError("Could not update section", err)
	}

	return course, nil
}

func (cs *UpdateSectionUseCase) getAvaliationUrl(url string) (*value_objects.Url, error) {
	if url == "" {
		return nil, nil
	}

	discord, err := value_objects.NewUrl(url)
	if err != nil {
		return nil, err
	}

	return discord, nil
}
