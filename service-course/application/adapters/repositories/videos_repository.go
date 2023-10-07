package repositories

import (
	"context"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/ports"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type VideosRepository struct {
	db *mongo.Database
}

type VideoModel struct {
	Id        string    `bson:"_id"`
	Type      string    `bson:"type"`
	TmpUrl    string    `bson:"url"`
	Processed bool      `bson:"processed"`
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
		Processed: video.Processed(),
		Duration:  video.Duration(),
		Size:      video.Size(),
		CreatedAt: video.CreatedAt(),
		UpdatedAt: video.UpdatedAt(),
	}

	_, err := collection.InsertOne(context.Background(), insertModel)

	return err
}
