package mappers

import (
	"github.com/matheusvmallmann/plataforma-ead/service-course/application/adapters/repositories/models"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
)

func CourseModelToEntityMap(Model models.CourseModel) *entities.Course {
	course := entities.NewCourseComplete(
		Model.Id,
		Model.Name,
		Model.Description,
		nil,
		Model.InstructorId,
		Model.Visible,
		Model.CreatedAt,
		Model.UpdatedAt,
	)

	for _, section := range Model.Sections {
		sectionEnity := entities.NewCompleteSection(
			entities.NewCompleteSectionData{
				Id:          section.Id,
				Name:        section.Name,
				Description: section.Description,
				CreatedAt:   section.CreatedAt,
				UpdatedAt:   section.UpdatedAt,
			},
		)

		for _, item := range section.Itens {
			itemEntity := entities.NewCourseItemComplete(
				item.Id,
				item.Name,
				item.Description,
				section.Id,
				item.Type,
				item.VideoId,
				item.CreatedAt,
				item.UpdatedAt,
			)
			sectionEnity.AddItem(itemEntity)
		}

		course.AddSection(sectionEnity)
	}

	return course
}
