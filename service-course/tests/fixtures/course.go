package fixtures

import (
	"time"

	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
)

var CursoCompleto = bson.M{
	"_id":          "course_id",
	"name":         "Teste",
	"description":  "Teste teste teste dsfdsfd dsfsdfdf",
	"instructorId": "user_id_1",
	"visible":      false,
	"sections": []bson.M{
		{
			"_id":         "section_id_1",
			"name":        "Seção 1 do curso",
			"description": "Descrição da seção",
			"order":       1,
			"itens": []bson.M{
				{
					"_id":         "item_id_1",
					"title":       "Item 1",
					"description": "Item 1 description",
					"videoId":     "video_id_1",
					"order":       1,
				},
				{
					"_id":         "item_id_2",
					"title":       "Item 2",
					"description": "Item 2 description",
					"videoId":     "video_id_2",
					"order":       2,
				},
			},
		},
		{
			"_id":         "section_id_2",
			"name":        "Seção 2 do curso",
			"description": "Descrição da seção",
			"order":       2,
			"itens":       nil,
		},
	},
}

func NewCourseFixture() *entities.Course {

	course := entities.NewCourseComplete(
		"course_id_1",
		"Sample Course",
		"This is a sample course description",
		nil,
		"user_id",
		true,
		time.Now(),
		time.Now(),
	)

	section1 := entities.NewCompleteSection(
		entities.NewCompleteSectionData{
			Id:          "section_id_1",
			Name:        "Section 1",
			Description: "Description for Section 1",
			CourseId:    "course_id_1",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	)

	section2 := entities.NewCompleteSection(
		entities.NewCompleteSectionData{
			Id:          "section_id_2",
			Name:        "Section 2",
			Description: "Description for Section 3",
			CourseId:    "course_id_1",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	)

	section3 := entities.NewCompleteSection(
		entities.NewCompleteSectionData{
			Id:          "section_id_2",
			Name:        "Section 3",
			Description: "Description for Section 3",
			CourseId:    "course_id_1",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	)

	item1 := entities.NewCourseItemComplete(
		"item_id_1",
		"Item 1",
		"Item 1 description",
		"section_id_1",
		"video",
		"",
		time.Now(),
		time.Now(),
	)

	section1.AddItem(item1)

	course.AddSection(section1)
	course.AddSection(section2)
	course.AddSection(section3)
	return course
}
