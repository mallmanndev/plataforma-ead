package repositories

import (
	"context"
	"errors"

	"github.com/matheusvmallmann/plataforma-ead/backend/modules/users/aplication/adapters/repositories/models"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/users/domain/entities"
	value_objects "github.com/matheusvmallmann/plataforma-ead/backend/modules/users/domain/value-objects"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersRepository struct {
	collection *mongo.Collection
}

func NewUsersRepository(db *mongo.Database) *UsersRepository {
	collection := db.Collection("users")
	return &UsersRepository{collection: collection}
}

func (r *UsersRepository) Create(user *entities.User) error {
	_, err := r.collection.InsertOne(context.Background(), models.User{
		ID:        user.Id,
		Name:      user.Name,
		Email:     user.Email.Email,
		Phone:     user.Phone.Phone,
		Password:  user.GetPassword(),
		Type:      user.Type.Id,
		CreatedAt: user.CreatedAt,
	})
	return err
}

func (r *UsersRepository) Update(user *entities.User) error {
	filter := bson.M{"_id": user.Id}
	update := bson.M{"$set": models.User{
		ID:        user.Id,
		Name:      user.Name,
		Email:     user.Email.Email,
		Phone:     user.Phone.Phone,
		Password:  user.GetPassword(),
		Type:      user.Type.Id,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}}
	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	return err
}

func (r *UsersRepository) Delete(id string) error {
	_, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}

func (r *UsersRepository) FindByEmail(email *value_objects.EmailAddress) (*entities.User, error) {
	filter := bson.M{"email": email.Email}
	model := &models.User{}
	err := r.collection.FindOne(context.Background(), filter).Decode(model)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil // user not found
		}
		return nil, err
	}

	user := entities.NewUserFromRepository(model)

	return user, nil
}

func (r *UsersRepository) FindById(id string) (*entities.User, error) {
	model := &models.User{}
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(model)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil // user not found
		}
		return nil, err
	}

	user := entities.NewUserFromRepository(model)

	return user, nil
}
