package usecases

import (
	errs "github.com/matheusvmallmann/plataforma-ead/service-course/application/errors"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/ports"
)

type CreateItem struct {
	coursesRepo ports.CourseRepository
	videosRepo  ports.VideosRepository
}

type CreateItemInput struct {
	SectionId   string
	UserId      string
	Title       string
	Description string
	VideoId     string
}

func NewCreateItem(coursesRepo ports.CourseRepository, videosRepo ports.VideosRepository) *CreateItem {
	return &CreateItem{coursesRepo, videosRepo}
}

func (ci *CreateItem) Execute(Data CreateItemInput) error {
	course, _ := ci.coursesRepo.FindBySectionId(Data.SectionId)
	if course == nil {
		return errs.NewNotFoundError("Course")
	}
	if course.InstructorID() != Data.UserId {
		return errs.NewPermissionDeniedError("create item")
	}

	video, _ := ci.videosRepo.Find(Data.VideoId)
	if video == nil {
		return errs.NewNotFoundError("Video")
	}
	if video.UserId() != Data.UserId {
		return errs.NewPermissionDeniedError("create item")
	}

	section := course.FindSection(Data.SectionId)

	itemType := "video"
	item := entities.NewCourseItem(
		Data.Title,
		Data.Description,
		Data.SectionId,
		itemType,
		Data.VideoId,
	)
	section.AddItem(item)

	if err := ci.coursesRepo.Update(course); err != nil {
		return err
	}
	return nil
}
