package usecases

import (
	errs "github.com/matheusvmallmann/plataforma-ead/service-course/application/errors"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/ports"
)

type DeleteCourseUseCase struct {
	courseRepository ports.CourseRepository
}

func NewDeleteCourseUseCase(CourseRepository ports.CourseRepository) *DeleteCourseUseCase {
	return &DeleteCourseUseCase{CourseRepository}
}

type DeleteCourseUseCaseDataDTO struct {
	Id     string
	UserId string
}

func (dc *DeleteCourseUseCase) Execute(Data DeleteCourseUseCaseDataDTO) error {
	findCourse, _ := dc.courseRepository.FindById(Data.Id)
	if findCourse == nil {
		return errs.NewNotFoundError("Course")
	}
	if findCourse.UserId() != Data.UserId {
		return errs.NewPermissionDeniedError("delete course")
	}

	if err := dc.courseRepository.Delete(Data.Id); err != nil {
		return errs.NewDeleteCourseUseCaseError("Could not delete course", err)
	}

	return nil
}
