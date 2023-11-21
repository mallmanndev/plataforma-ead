package usecases

import (
	errs "github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/errors"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/ports"
	value_objects "github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/value-objects"
)

type UpdateCourseUseCase struct {
	courseRepository ports.CourseRepository
}

func NewUpdateCourseUseCase(
	CourseRepository ports.CourseRepository,
) *UpdateCourseUseCase {
	return &UpdateCourseUseCase{
		CourseRepository,
	}
}

type UpdateCourseUseCaseDTO struct {
	Id          string
	Name        string
	Description string
	DiscordUrl  string
	UserId      string
}

func (uc *UpdateCourseUseCase) Execute(Data UpdateCourseUseCaseDTO) (*entities.Course, error) {
	course, _ := uc.courseRepository.FindById(Data.Id)
	if course == nil {
		return nil, errs.NewNotFoundError("Course")
	}

	if course.UserId() != Data.UserId {
		return nil, errs.NewPermissionDeniedError("update course")
	}

	discord, err := uc.getDiscordUrl(Data.DiscordUrl)
	if err != nil {
		return nil, err
	}

	if err := course.Update(Data.Name, Data.Description, discord); err != nil {
		return nil, err
	}

	if err := uc.courseRepository.Update(course); err != nil {
		return nil, errs.NewUpdateCourseUseCaseError("Could not update course", err)
	}

	return course, nil
}

func (cc *UpdateCourseUseCase) getDiscordUrl(url string) (*value_objects.Url, error) {
	if url == "" {
		return nil, nil
	}

	discord, err := value_objects.NewUrl(url)
	if err != nil {
		return nil, err
	}

	return discord, nil
}
