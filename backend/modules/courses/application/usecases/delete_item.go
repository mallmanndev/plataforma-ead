package usecases

import (
	errs "github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/errors"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/ports"
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
	if course == nil || course.UserId() != Data.UserId {
		return nil, errs.NewNotFoundError("Item")
	}

	item, section := course.FindItem(Data.Id)

	section.RemoveItem(item.Id())
	if err := ui.coursesRepo.Update(course); err != nil {
		return nil, err
	}

	return course, nil
}
