package usecases

import (
	errs "github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/errors"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/ports"
)

type MakeCourseInvisible struct {
	CourseRepository ports.CourseRepository
}

func NewMakeCourseInvisible(courseRepository ports.CourseRepository) *MakeCourseInvisible {
	return &MakeCourseInvisible{CourseRepository: courseRepository}
}

func (m MakeCourseInvisible) Execute(id string, userId string) error {
	course, err := m.CourseRepository.FindById(id)
	if err != nil {
		return err
	}

	if course.UserId() != userId {
		return errs.NewPermissionDeniedError("cahange visibility of this course")
	}

	course.MakeInvisible()

	return m.CourseRepository.Update(course)
}
