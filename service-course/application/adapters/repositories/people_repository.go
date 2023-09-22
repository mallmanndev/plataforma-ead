package repositories

import (
	"context"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PeopleRepository struct {
	db *mongo.Client
}

func NewPeopleRepository(db *mongo.Client) *PeopleRepository {
	return &PeopleRepository{db}
}

type PeopleModel struct {
	Id       string `bson:"_id"`
	Name     string `bson:"name"`
	Type     string `bson:"type"`
	PhotoUrl string `bson:"photoUrl"`
}

func (pr *PeopleRepository) Upsert(People *entities.People) error {
	collection := pr.db.Database("service-courses").Collection("people")
	var photoUrl string = ""
	if People.Photo() != nil {
		photoUrl = People.Photo().Url()
	}
	insertModel := PeopleModel{
		Id:       People.Id(),
		Name:     People.Name(),
		Type:     People.GetType(),
		PhotoUrl: photoUrl,
	}
	opts := options.Update().SetUpsert(true)
	_, err := collection.UpdateOne(
		context.Background(),
		bson.M{"_id": People.Id()},
		bson.M{"$set": insertModel},
		opts)
	return err
}

func (pr *PeopleRepository) FindById(Id string) (*entities.People, error) {
	return nil, nil
}
