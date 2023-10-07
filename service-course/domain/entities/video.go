package entities

import (
	errs "github.com/matheusvmallmann/plataforma-ead/service-course/application/errors"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/apptimer"
	"time"
)

type Video struct {
	timer     apptimer.Timer
	id        string
	videoType string
	tmpUrl    string
	processed bool
	duration  float32
	size      int64
	createdAt time.Time
	updatedAt time.Time
}

func NewVideo(Timer apptimer.Timer, Id string, TmpUrl string, Type string, Size int64) (*Video, error) {
	video := &Video{
		timer:     Timer,
		id:        Id,
		tmpUrl:    TmpUrl,
		videoType: Type,
		processed: false,
		size:      Size,
		createdAt: Timer.Now(),
	}

	if err := video.Validate(); err != nil {
		return nil, err
	}

	return video, nil
}

func (v *Video) Validate() error {
	if v.videoType != "mp4" {
		return errs.NewInvalidAttributeError(
			"Video",
			"type",
			"must be mp4.",
		)
	}

	return nil
}

func (v *Video) Id() string {
	return v.id
}

func (v *Video) Type() string {
	return v.videoType
}

func (v *Video) TmpUrl() string {
	return v.tmpUrl
}

func (v *Video) Processed() bool {
	return v.processed
}

func (v *Video) Duration() float32 {
	return v.duration
}

func (v *Video) Size() int64 {
	return v.size
}

func (v *Video) CreatedAt() time.Time {
	return v.createdAt
}

func (v *Video) UpdatedAt() time.Time {
	return v.updatedAt
}
