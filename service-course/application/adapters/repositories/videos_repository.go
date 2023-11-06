package repositories

import (
	"context"
	"errors"

	"github.com/matheusvmallmann/plataforma-ead/service-course/application/adapters/repositories/mappers"
	"github.com/matheusvmallmann/plataforma-ead/service-course/application/adapters/repositories/models"
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

func NewVideosRepository(db *mongo.Database) ports.VideosRepository {
	return &VideosRepository{db: db}
}

func (vr *VideosRepository) Create(video *entities.Video) error {
	collection := vr.db.Collection("videos")

	insertModel := models.VideoModel{
		Id:        video.Id(),
		Type:      video.Type(),
		TmpUrl:    video.TmpUrl(),
		Status:    video.Status(),
		Duration:  video.Duration(),
		Size:      video.Size(),
		Url:       video.Url(),
		CreatedAt: video.CreatedAt(),
		UpdatedAt: video.UpdatedAt(),
		UserId:    video.UserId(),
	}

	_, err := collection.InsertOne(context.Background(), insertModel)

	return err
}

func (vr *VideosRepository) Find(Id string) (*entities.Video, error) {
	collection := vr.db.Collection("videos")
	var video models.VideoModel
	err := collection.FindOne(context.Background(), bson.M{"_id": Id}).Decode(&video)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return mappers.VideoModelToVideoEntity(video), nil
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
		var videoModel models.VideoModel
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
			videoModel.UserId,
			videoModel.CreatedAt,
			videoModel.UpdatedAt,
			videoModel.Url,
		)
		videos = append(videos, video)
	}
	return videos, nil
}

func (vr *VideosRepository) Update(video *entities.Video) error {
	collection := vr.db.Collection("videos")

	filter := bson.M{"_id": video.Id()}

	var resolutions []models.VideoResolution
	for _, res := range video.GetResolutions() {
		resolutions = append(resolutions, models.VideoResolution{
			Resolution: res.Resolution,
			URL:        res.URL,
		})
	}

	update := bson.M{"$set": bson.M{
		"type":        video.Type(),
		"tmp":         video.TmpUrl(),
		"status":      video.Status(),
		"duration":    video.Duration(),
		"size":        video.Size(),
		"updatedAt":   video.UpdatedAt(),
		"resolutions": resolutions,
		"url":         video.Url(),
	}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	return err
}

func (vr *VideosRepository) Delete(id string) error {
	return errors.New("not implemented")
}
