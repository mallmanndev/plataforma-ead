package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/apptimer"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var timer = apptimer.NewAppTimer()

type VideosRepository struct {
	db *mongo.Database
}

type VideoModel struct {
	Id        string    `bson:"_id"`
	Type      string    `bson:"type"`
	TmpUrl    string    `bson:"url"`
	Status    string    `bson:"status"`
	Duration  float32   `bson:"duration"`
	Size      int64     `bson:"size"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

func NewVideosRepository(db *mongo.Database) ports.VideosRepository {
	return &VideosRepository{db: db}
}

func (vr *VideosRepository) Create(video *entities.Video) error {
	collection := vr.db.Collection("videos")

	insertModel := VideoModel{
		Id:        video.Id(),
		Type:      video.Type(),
		TmpUrl:    video.TmpUrl(),
		Status:    video.Status(),
		Duration:  video.Duration(),
		Size:      video.Size(),
		CreatedAt: video.CreatedAt(),
		UpdatedAt: video.UpdatedAt(),
	}

	_, err := collection.InsertOne(context.Background(), insertModel)

	return err
}

func (vr *VideosRepository) Find(Id string) (*entities.Video, error) {
	return nil, errors.New("not implemented")
}

func (vr *VideosRepository) Get(filters ports.GetFilters) ([]*entities.Video, error) {
	var videos []*entities.Video
	ctx := context.Background()
	collection := vr.db.Collection("videos")
	cursor, err := collection.Find(ctx, bson.D{{Key: "status", Value: filters.Status}})
	if err != nil {
		return videos, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var videoModel VideoModel
		err := cursor.Decode(&videoModel)
		if err != nil {
			return nil, err
		}
		video := entities.NewCompleteVideo(
			timer, videoModel.Id,
			videoModel.Type,
			videoModel.TmpUrl,
			videoModel.Status,
			videoModel.Duration,
			videoModel.Size,
			videoModel.CreatedAt,
			videoModel.UpdatedAt,
		)
		videos = append(videos, video)
	}
	return videos, nil
}

func (vr *VideosRepository) Update(video *entities.Video) error {
	collection := vr.db.Collection("videos")

	filter := bson.M{"_id": video.Id()}

	update := bson.M{"$set": bson.M{
		"type":      video.Type(),
		"url":       video.TmpUrl(),
		"status":    video.Status(),
		"duration":  video.Duration(),
		"size":      video.Size(),
		"updatedAt": video.UpdatedAt(),
	}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	return err
}

func (vr *VideosRepository) Delete(id string) error {
	return errors.New("not implemented")
}
