package fixtures

import (
	"time"

	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/apptimer"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
)

func NewVideo() *entities.Video {
	return entities.NewCompleteVideo(
		apptimer.NewFakeTimer(time.Date(2023, time.October, 6, 12, 0, 0, 0, time.UTC)),
		"video_id",
		"mp4",
		"",
		"processed",
		10000,
		10000,
		"user_id",
		time.Date(2023, time.October, 6, 12, 0, 0, 0, time.UTC),
		time.Date(2023, time.October, 6, 12, 0, 0, 0, time.UTC),
		"/test/video",
	)
}

func NewUser2Video() *entities.Video {
	return entities.NewCompleteVideo(
		apptimer.NewFakeTimer(time.Date(2023, time.October, 6, 12, 0, 0, 0, time.UTC)),
		"video_id",
		"mp4",
		"",
		"processed",
		10000,
		10000,
		"user_id_2",
		time.Date(2023, time.October, 6, 12, 0, 0, 0, time.UTC),
		time.Date(2023, time.October, 6, 12, 0, 0, 0, time.UTC),
		"/test/video",
	)
}

var VideosBson = []interface{}{
	bson.M{
		"_id":      "video_id_1",
		"type":     "mp4",
		"url":      "/videos/tmp/video_id_1.mp4",
		"status":   "success",
		"duration": 0,
		"userId":   "user_id_1",
		"size":     408532919,
		"resolutions": []bson.M{
			{
				"resolution": "480",
				"url":        "/videos/video_id_1/480",
			},
			{
				"resolution": "720",
				"url":        "/videos/video_id_1/720",
			},
			{
				"resolution": "1080",
				"url":        "/videos/video_id_1/1080",
			},
		},
	},
	bson.M{
		"_id":      "video_id_2",
		"type":     "mp4",
		"url":      "/videos/tmp/video_id_2.mp4",
		"status":   "success",
		"duration": 0,
		"userId":   "user_id_2",
		"size":     408532919,
		"resolutions": []bson.M{
			{
				"resolution": "480",
				"url":        "/videos/video_id_2/480",
			},
			{
				"resolution": "720",
				"url":        "/videos/video_id_2/720",
			},
			{
				"resolution": "1080",
				"url":        "/videos/video_id_2/1080",
			},
		},
	},
}
