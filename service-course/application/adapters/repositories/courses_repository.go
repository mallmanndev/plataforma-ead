package repositories

import (
	"context"
	"log"

	"github.com/matheusvmallmann/plataforma-ead/service-course/application/adapters/repositories/mappers"
	"github.com/matheusvmallmann/plataforma-ead/service-course/application/adapters/repositories/models"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CoursesRepositories struct {
	collection *mongo.Collection
}

func NewCourseRepositories(Db *mongo.Database) ports.CourseRepository {
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

func (cr *CoursesRepositories) FindBySectionId(Id string) (*entities.Course, error) {
	courseModel := models.CourseModel{}

	err := cr.collection.FindOne(context.Background(), bson.M{"sections._id": Id}).Decode(&courseModel)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return mappers.CourseModelToEntityMap(courseModel), nil
}

func (cr *CoursesRepositories) FindByItemId(Id string) (*entities.Course, error) {
	courseModel := models.CourseModel{}

	err := cr.collection.FindOne(context.Background(), bson.M{"sections.itens._id": Id}).Decode(&courseModel)
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
		Id:          Course.Id(),
		Name:        Course.Name(),
		Description: Course.Description(),
		UserId:      Course.UserId(),
		Sections:    []models.CourseSectionModel{},
		CreatedAt:   Course.CreatedAt(),
	})
	return err
}

func (cr *CoursesRepositories) Update(Course *entities.Course) error {
	filter := bson.M{"_id": Course.Id()}

	var sections []models.CourseSectionModel

	for _, section := range Course.Sections() {
		var itens []models.CourseItemModel

		for _, item := range section.Itens() {
			newItem := models.CourseItemModel{
				Id:          item.Id(),
				Title:       item.Title(),
				Description: item.Description(),
				Type:        item.Type(),
				VideoId:     item.VideoId(),
				Order:       item.Order(),
				CreatedAt:   item.CreatedAt(),
				UpdatedAt:   item.UpdatedAt(),
			}
			itens = append(itens, newItem)
		}

		newSection := models.CourseSectionModel{
			Id:          section.Id(),
			Name:        section.Name(),
			Description: section.Description(),
			CreatedAt:   section.CreatedAt(),
			UpdatedAt:   section.UpdatedAt(),
			Itens:       itens,
		}
		sections = append(sections, newSection)
	}

	update := bson.M{
		"$set": bson.M{
			"name":        Course.Name(),
			"description": Course.Description(),
			"sections":    sections,
			"updatedAt":   Course.UpdatedAt(),
		},
	}
	_, err := cr.collection.UpdateOne(context.Background(), filter, update)
	return err
}

func (cr *CoursesRepositories) Delete(Id string) error {
	_, err := cr.collection.DeleteOne(context.Background(), bson.M{"_id": Id})
	return err
}

func (cr *CoursesRepositories) Get(Filters ports.GetCourseFilters) ([]*entities.Course, error) {
	filter := bson.M{}
	if Filters.UserId != "" {
		filter["userId"] = Filters.UserId
	}
	if Filters.Id != "" {
		filter["_id"] = Filters.Id
	}
	if Filters.Visible {
		filter["visible"] = true
	}

	cursor, err := cr.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	var courses []*entities.Course

	for cursor.Next(context.Background()) {
		courseModel := models.CourseModel{}
		if err := cursor.Decode(&courseModel); err != nil {
			log.Fatal(err)
		}

		courses = append(courses, mappers.CourseModelToEntityMap(courseModel))
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return courses, nil
}
