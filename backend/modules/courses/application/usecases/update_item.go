package usecases

import (
	errs "github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/errors"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/ports"
)

type UpdateItem struct {
	coursesRepo ports.CourseRepository
}

type UpdateItemInput struct {
	Id          string
	UserId      string
	Title       string
	Description string
}

func NewUpdateItem(CoursesRepo ports.CourseRepository) *UpdateItem {
	return &UpdateItem{coursesRepo: CoursesRepo}
}

func (ui *UpdateItem) Execute(Data UpdateItemInput) (*entities.Course, error) {
	course, _ := ui.coursesRepo.FindByItemId(Data.Id)
	if course == nil || course.UserId() != Data.UserId {
		return nil, errs.NewNotFoundError("Item")
	}

	item, _ := course.FindItem(Data.Id)

	item.Update(Data.Title, Data.Description)
	if err := ui.coursesRepo.Update(course); err != nil {
		return nil, err
	}

	return course, nil
}
