package usecases

import (
	errs "github.com/matheusvmallmann/plataforma-ead/service-course/application/errors"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/ports"
)

type UpdateCourseUseCase struct {
	peopleRepository ports.PeopleRepository
	courseRepository ports.CourseRepository
}

func NewUpdateCourseUseCase(
	PeopleRepository ports.PeopleRepository,
	CourseRepository ports.CourseRepository,
) *UpdateCourseUseCase {
	return &UpdateCourseUseCase{
		PeopleRepository,
		CourseRepository,
	}
}

type UpdateCourseInstructorDTO struct {
	Id   string
	Name string
	Type string
}

type UpdateCourseUseCaseDTO struct {
	Id          string
	Name        string
	Description string
	Instructor  UpdateCourseInstructorDTO
}

func (uc *UpdateCourseUseCase) Execute(Data UpdateCourseUseCaseDTO) (*entities.Course, error) {
	people, err := entities.NewPeople(Data.Instructor.Id, Data.Instructor.Name, Data.Instructor.Type, nil)
	if err != nil {
		return nil, err
	}

	if err := uc.peopleRepository.Upsert(people); err != nil {
		return nil, errs.NewUpdateCourseUseCaseError("Could not insert or update people", err)
	}

	course, _ := uc.courseRepository.FindById(Data.Id)
	if course == nil {
		return nil, errs.NewDataNotFoundError("Course not found!")
	}

	if err := course.Update(Data.Name, Data.Description); err != nil {
		return nil, err
	}

	if err := uc.courseRepository.Update(course); err != nil {
		return nil, errs.NewUpdateCourseUseCaseError("Could not update course", err)
	}

	return course, nil
}
