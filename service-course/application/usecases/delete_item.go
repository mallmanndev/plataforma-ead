package usecases

import (
	errs "github.com/matheusvmallmann/plataforma-ead/service-course/application/errors"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/ports"
)

type DeleteItem struct {
	coursesRepo ports.CourseRepository
}

type DeleteItemInput struct {
	Id     string
	UserId string
}

func NewDeleteItem(CoursesRepo ports.CourseRepository) *DeleteItem {
	return &DeleteItem{CoursesRepo}
}

func (ui *DeleteItem) Execute(Data DeleteItemInput) (*entities.Course, error) {
	course, _ := ui.coursesRepo.FindByItemId(Data.Id)
	if course == nil || course.InstructorID() != Data.UserId {
		return nil, errs.NewNotFoundError("Item")
	}

	item, section := course.FindItem(Data.Id)

	section.RemoveItem(item.Id())
	if err := ui.coursesRepo.Update(course); err != nil {
		return nil, err
	}

	return course, nil
}
