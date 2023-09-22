package repositories

import (
	"context"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type CoursesRepositories struct {
	db *mongo.Client
}

func NewCourseRepositories(Db *mongo.Client) *CoursesRepositories {
	return &CoursesRepositories{db: Db}
}

type CourseModel struct {
	Id           string    `bson:"_id"`
	Name         string    `bson:"name"`
	Description  string    `bson:"description"`
	InstructorId string    `bson:"instructorId"`
	Visible      bool      `bson:"visible"`
	Sections     []any     `bson:"sections"`
	CreatedAt    time.Time `bson:"createdAt"`
}

func (cr *CoursesRepositories) FindById(Id string) (*entities.Course, error) {
	return nil, nil
}

func (cr *CoursesRepositories) Create(Course *entities.Course) error {
	collection := cr.db.Database("service-courses").Collection("courses")
	_, err := collection.InsertOne(context.Background(), CourseModel{
		Id:           Course.Id(),
		Name:         Course.Name(),
		Description:  Course.Description(),
		InstructorId: Course.InstructorID(),
		Sections:     []any{},
		CreatedAt:    Course.CreatedAt(),
	})
	return err
}

func (cr *CoursesRepositories) Update(Course *entities.Course) error {
	return nil
}

func (cr *CoursesRepositories) Delete(Id string) error {
	return nil
}

func (cr *CoursesRepositories) AddSection(Section *entities.CourseSection) error {
	return nil
}

func (cr *CoursesRepositories) UpdateSection(Section *entities.CourseSection) error {
	return nil
}

func (cr *CoursesRepositories) RemoveSection(Id string) error {
	return nil
}
