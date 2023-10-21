package mappers

import (
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/service-course/pb"
)

func CourseEnitiyToGrpc(course *entities.Course) *pb.Course {
	var sections []*pb.CourseSection

	for _, section := range course.Sections() {
		var itens []*pb.CourseItem

		for _, item := range section.Itens() {
			grpcItem := &pb.CourseItem{
				Id:          item.Id(),
				Title:       item.Title(),
				Description: item.Description(),
				VideoId:     item.VideoId(),
				CreatedAt:   item.CreatedAt().String(),
				UpdatedAt:   item.UpdatedAt().String(),
			}
			itens = append(itens, grpcItem)
		}

		grpcSection := &pb.CourseSection{
			Id:          section.Id(),
			Name:        section.Name(),
			Description: section.Description(),
			CreatedAt:   section.CreatedAt().String(),
			UpdatedAt:   section.UpdatedAt().String(),
			Itens:       itens,
		}
		sections = append(sections, grpcSection)
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
