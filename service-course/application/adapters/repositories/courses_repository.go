package repositories

import (
	"context"
	"github.com/matheusvmallmann/plataforma-ead/service-course/application/adapters/repositories/mappers"
	"github.com/matheusvmallmann/plataforma-ead/service-course/application/adapters/repositories/models"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CoursesRepositories struct {
	collection *mongo.Collection
}

func NewCourseRepositories(Db *mongo.Database) *CoursesRepositories {
	collection := Db.Collection("courses")
	return &CoursesRepositories{collection}
}

func (cr *CoursesRepositories) FindById(Id string) (*entities.Course, error) {
	courseModel := models.CourseModel{}

	err := cr.collection.FindOne(context.Background(), bson.M{"_id": Id}).Decode(&courseModel)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return mappers.CourseModelToEntityMap(courseModel), nil
}

func (cr *CoursesRepositories) Create(Course *entities.Course) error {
	_, err := cr.collection.InsertOne(context.Background(), models.CourseModel{
		Id:           Course.Id(),
		Name:         Course.Name(),
		Description:  Course.Description(),
		InstructorId: Course.InstructorID(),
		Sections:     []models.CourseSectionModel{},
		CreatedAt:    Course.CreatedAt(),
	})
	return err
}

func (cr *CoursesRepositories) Update(Course *entities.Course) error {
	filter := bson.M{"_id": Course.Id()}

	var sections []models.CourseSectionModel

	for _, section := range Course.Sections() {
		newSection := models.CourseSectionModel{
			Id:          section.Id(),
			Name:        section.Name(),
			Description: section.Description(),
			CreatedAt:   section.CreatedAt(),
			UpdatedAt:   section.UpdatedAt(),
		}
		sections = append(sections, newSection)
	}

	update := bson.M{"$set": bson.M{
		"name":        Course.Name(),
		"description": Course.Description(),
		"sections":    sections,
		"updatedAt":   Course.UpdatedAt(),
	}}
	_, err := cr.collection.UpdateOne(context.Background(), filter, update)
	return err
}

func (cr *CoursesRepositories) Delete(Id string) error {
	_, err := cr.collection.DeleteOne(context.Background(), bson.M{"_id": Id})
	return err
}
