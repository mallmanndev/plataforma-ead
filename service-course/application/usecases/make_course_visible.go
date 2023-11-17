package usecases

import (
	errs "github.com/matheusvmallmann/plataforma-ead/service-course/application/errors"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/ports"
)

type MakeCourseVisible struct {
	CourseRepository ports.CourseRepository
}

func NewMakeCourseVisible(courseRepository ports.CourseRepository) *MakeCourseVisible {
	return &MakeCourseVisible{CourseRepository: courseRepository}
}

func (m MakeCourseVisible) Execute(id string, userId string) error {
	course, err := m.CourseRepository.FindById(id)
	if err != nil {
		return err
	}

	if course.UserId() != userId {
		return errs.NewPermissionDeniedError("cahange visibility of this course")
	}

	if err := course.MakeVisible(); err != nil {
		return err
	}

	return m.CourseRepository.Update(course)
}
