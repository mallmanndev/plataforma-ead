package usecases

import (
	errs "github.com/matheusvmallmann/plataforma-ead/service-course/application/errors"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/ports"
)

type DeleteSectionUseCase struct {
	coursesRepository ports.CourseRepository
}

func NewDeleteSectionUseCase(CoursesRepository ports.CourseRepository) *DeleteSectionUseCase {
	return &DeleteSectionUseCase{CoursesRepository}
}

type DeleteSectionDTO struct {
	UserId    string
	SectionId string
}

func (ds *DeleteSectionUseCase) Execute(Data DeleteSectionDTO) error {
	course, _ := ds.coursesRepository.FindBySectionId(Data.SectionId)
	if course == nil {
		return errs.NewDeleteSectionUseCaseError("Course not found", nil)
	}
	if course.UserId() != Data.UserId {
		return errs.NewPermissionDeniedError("update section")
	}

	if err := course.RemoveSection(Data.SectionId); err != nil {
		return errs.NewNotFoundError("Section")
	}

	if err := ds.coursesRepository.Update(course); err != nil {
		return errs.NewDeleteSectionUseCaseError("Could not delete section", err)
	}
	return nil
}
