package mappers

import (
	"fmt"

	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/adapters/repositories/models"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/entities"
	value_objects "github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/value-objects"
)

func CourseModelToEntityMap(Model models.CourseModel) *entities.Course {
	course := entities.NewCourseComplete(
		Model.Id,
		Model.Name,
		Model.Description,
		nil,
		Model.UserId,
		Model.Visible,
		Model.CreatedAt,
		Model.UpdatedAt,
		Model.DiscordUrl,
	)

	for _, section := range Model.Sections {

		var avaliation *value_objects.Url
		if section.AvaliationUrl != "" {
			fmt.Println(section.AvaliationUrl)
			avaliation, _ = value_objects.NewUrl(section.AvaliationUrl)
		}

		sectionEnity := entities.NewCompleteSection(
			entities.NewCompleteSectionData{
				Id:          section.Id,
				Name:        section.Name,
				Description: section.Description,
				CreatedAt:   section.CreatedAt,
				UpdatedAt:   section.UpdatedAt,
				Avaliation:  avaliation,
			},
		)

		for _, item := range section.Itens {
			itemEntity := entities.NewCourseItemComplete(
				item.Id,
				item.Title,
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
