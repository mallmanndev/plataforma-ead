package entities

import (
	"sort"
	"time"

	errs "github.com/matheusvmallmann/plataforma-ead/service-course/application/errors"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/apptimer"
)

type VideoResolution struct {
	Resolution         int
	CompleteResolution string
	URL                string
}

type Video struct {
	timer       apptimer.Timer
	id          string
	videoType   string
	tmpUrl      string
	status      string
	duration    float32
	size        int64
	createdAt   time.Time
	updatedAt   time.Time
	resolutions []VideoResolution
}

func NewVideo(Timer apptimer.Timer, Id string, TmpUrl string, Type string, Size int64) (*Video, error) {
	video := &Video{
		timer:     Timer,
		id:        Id,
		tmpUrl:    TmpUrl,
		videoType: Type,
		status:    "pending",
		size:      Size,
		createdAt: Timer.Now(),
	}
	if err := video.Validate(); err != nil {
		return nil, err
	}
	return video, nil
}

func NewCompleteVideo(
	timer apptimer.Timer,
	id string,
	videoType string,
	tmpUrl string,
	status string,
	duration float32,
	size int64,
	createdAt time.Time,
	updatedAt time.Time,
) *Video {
	return &Video{
		timer:     timer,
		id:        id,
		videoType: videoType,
		tmpUrl:    tmpUrl,
		status:    status,
		duration:  duration,
		size:      size,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
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

func (v *Video) SetStatus(Status string) *Video {
	v.status = Status
	return v
}

func (v *Video) AddResolution(Resolution VideoResolution) {
	res := append(v.resolutions, Resolution)
	sort.Slice(res, func(i int, j int) bool {
		return res[i].Resolution < res[j].Resolution
	})
	v.resolutions = res
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

func (v *Video) Status() string {
	return v.status
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

func (v *Video) GetResolutions() []VideoResolution {
	return v.resolutions
}
