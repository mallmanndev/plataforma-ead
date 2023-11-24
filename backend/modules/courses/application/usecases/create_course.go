package usecases

import (
	errs "github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/errors"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/ports"
	value_objects "github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/value-objects"
)

type CreateCourseUseCase struct {
	courseRepository ports.CourseRepository
}

func NewCreateCourseUseCase(
	CourseRepository ports.CourseRepository,
) *CreateCourseUseCase {
	return &CreateCourseUseCase{
		courseRepository: CourseRepository,
	}
}

type CreateCourseUseCaseDTO struct {
	Name        string
	Description string
	UserId      string
	DiscordUrl  string
}

func (cc *CreateCourseUseCase) Execute(Data CreateCourseUseCaseDTO) (*entities.Course, error) {
	discord, err := cc.getDiscordUrl(Data.DiscordUrl)
	if err != nil {
		return nil, err
	}

	course, err := entities.NewCourse(
		Data.Name,
		Data.Description,
		nil,
		Data.UserId,
		discord,
	)
	if err != nil {
		return nil, err
	}

	if err := cc.courseRepository.Create(course); err != nil {
		return nil, errs.NewCreateUserUseCaseError("Could not create course", err)
	}

	return course, nil
}

func (cc *CreateCourseUseCase) getDiscordUrl(url string) (*value_objects.Url, error) {
	if url == "" {
		return nil, nil
	}

	discord, err := value_objects.NewUrl(url)
	if err != nil {
		return nil, err
	}

	return discord, nil
}
