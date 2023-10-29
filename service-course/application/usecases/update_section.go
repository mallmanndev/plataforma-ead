package usecases

import (
	"log"

	errs "github.com/matheusvmallmann/plataforma-ead/service-course/application/errors"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/ports"
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
	if course.InstructorID() != Data.UserId {
		return nil, errs.NewPermissionDeniedError("update section")
	}

	section := course.FindSection(Data.SectionId)

	if err := section.Update(Data.Name, Data.Description); err != nil {
		return nil, err
	}

	if err := cs.coursesRepository.Update(course); err != nil {
		return nil, errs.NewUpdateSectionUseCaseError("Could not update section", err)
	}

	log.Println(course)
	return course, nil
}
