package fixtures

import (
	"time"

	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
)

var CursoCompleto = bson.M{
	"_id":          "3d515009-56eb-4ed0-aea5-182bd783085e",
	"name":         "Teste",
	"description":  "Teste teste teste dsfdsfd dsfsdfdf",
	"instructorId": "9111bffd-73d9-49d8-b32c-48353674dc06",
	"visible":      false,
	"sections": []bson.M{
		{
			"_id":         "3d515009-56eb-4ed0-aea5-182bd783085e",
			"name":        "Seção 1 do curso",
			"description": "Descrição da seção",
			"itens":       nil,
		},
		{
			"_id":         "3d515009-56eb-4ed0-aea5-182bd783ewfwe085e",
			"name":        "Seção 2 do curso",
			"description": "Descrição da seção",
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

	course.AddSection(section1)
	course.AddSection(section2)
	course.AddSection(section3)
	return course
}
