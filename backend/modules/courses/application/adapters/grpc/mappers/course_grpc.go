package mappers

import (
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/backend/pb"
)

func CourseEnitiyToGrpc(course *entities.Course) *pb.Course {
	var sections []*pb.CourseSection

	for _, section := range course.Sections() {
		sections = append(sections, SectionEntityToGrpc(section))
	}

	return &pb.Course{
		Id:          course.Id(),
		Name:        course.Name(),
		Description: course.Description(),
		Visible:     course.IsVisible(),
		Sections:    sections,
		CreatedAt:   course.CreatedAt().String(),
		UpdatedAt:   course.UpdatedAt().String(),
	}
}

func SectionEntityToGrpc(section *entities.CourseSection) *pb.CourseSection {
	var itens []*pb.CourseItem

	for _, item := range section.Itens() {
		itens = append(itens, ItemEntityToGrpc(item))
	}

	return &pb.CourseSection{
		Id:          section.Id(),
		Name:        section.Name(),
		Description: section.Description(),
		CreatedAt:   section.CreatedAt().String(),
		UpdatedAt:   section.UpdatedAt().String(),
		Itens:       itens,
	}
}

func ItemEntityToGrpc(item *entities.CourseItem) *pb.CourseItem {
	return &pb.CourseItem{
		Id:          item.Id(),
		Title:       item.Title(),
		Description: item.Description(),
		VideoId:     item.VideoId(),
		CreatedAt:   item.CreatedAt().String(),
		UpdatedAt:   item.UpdatedAt().String(),
	}
}
