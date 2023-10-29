package fixtures

import (
	"time"

	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/apptimer"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
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
