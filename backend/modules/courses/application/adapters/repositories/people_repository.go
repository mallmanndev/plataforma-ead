package repositories

import (
	"context"

	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PeopleRepository struct {
	db *mongo.Database
}

func NewPeopleRepository(db *mongo.Database) *PeopleRepository {
	return &PeopleRepository{db}
}

type PeopleModel struct {
	Id       string `bson:"_id"`
	Name     string `bson:"name"`
	Type     string `bson:"type"`
	PhotoUrl string `bson:"photoUrl"`
}

func (pr *PeopleRepository) Upsert(People *entities.People) error {
	collection := pr.db.Collection("people")
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
