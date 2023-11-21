package usecases

import (
	errs "github.com/matheusvmallmann/plataforma-ead/service-course/application/errors"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/ports"
	value_objects "github.com/matheusvmallmann/plataforma-ead/service-course/domain/value-objects"
)

type CreateCourseUseCase struct {
	peopleRepository ports.PeopleRepository
	courseRepository ports.CourseRepository
}

func NewCreateCourseUseCase(
	PeopleRepository ports.PeopleRepository,
	CourseRepository ports.CourseRepository,
) *CreateCourseUseCase {
	return &CreateCourseUseCase{
		peopleRepository: PeopleRepository,
		courseRepository: CourseRepository,
	}
}

type CreateCourseInstructorDTO struct {
	Id   string
	Name string
	Type string
}

type CreateCourseUseCaseDTO struct {
	Name        string
	Description string
	Instructor  CreateCourseInstructorDTO
	DiscordUrl  string
}

func (cc *CreateCourseUseCase) Execute(Data CreateCourseUseCaseDTO) (*entities.Course, error) {
	people, err := entities.NewPeople(
		Data.Instructor.Id,
		Data.Instructor.Name,
		Data.Instructor.Type,
		nil,
	)
	if err != nil {
		return nil, err
	}

	if err := cc.peopleRepository.Upsert(people); err != nil {
		return nil, errs.NewCreateUserUseCaseError("Could not insert or update people", err)
	}

	discord, err := cc.getDiscordUrl(Data.DiscordUrl)
	if err != nil {
		return nil, err
	}

	course, err := entities.NewCourse(
		Data.Name,
		Data.Description,
		nil,
		Data.Instructor.Id,
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
