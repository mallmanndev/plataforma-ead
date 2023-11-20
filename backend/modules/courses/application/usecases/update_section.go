package usecases

import (
	errs "github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/errors"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/ports"
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
	UserId      string
	SectionId   string
	Name        string
	Description string
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

	if err := section.Update(Data.Name, Data.Description); err != nil {
		return nil, err
	}

	if err := cs.coursesRepository.Update(course); err != nil {
		return nil, errs.NewUpdateSectionUseCaseError("Could not update section", err)
	}

	return course, nil
}
