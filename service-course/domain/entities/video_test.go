package entities_test

import (
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/apptimer"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewVideo(t *testing.T) {

	t.Run("Should return a error when type is not mp4", func(t *testing.T) {
		expectedTime := time.Date(2023, time.October, 6, 12, 0, 0, 0, time.UTC)
		fakeTimer := apptimer.NewFakeTimer(expectedTime)

		// WHEN
		video, err := entities.NewVideo(fakeTimer, "1", "url", "avi", 123)

		// THEN
		assert.Nil(t, video)
		if assert.Error(t, err) {
			assert.ErrorContains(t, err, "[Video] Invalid 'type': must be mp4.")
		}
	})

	t.Run("When video is valid, return a video struct", func(t *testing.T) {
		// GIVEN
		expectedTime := time.Date(2023, time.October, 6, 12, 0, 0, 0, time.UTC)
		fakeTimer := apptimer.NewFakeTimer(expectedTime)

		// WHEN
		video, err := entities.NewVideo(fakeTimer, "1", "url", "mp4", 123)

		// THEN
		assert.Nil(t, err)
		if assert.NotNil(t, video) {
			assert.Equal(t, "1", video.Id())
			assert.Equal(t, "url", video.TmpUrl())
			assert.Equal(t, "mp4", video.Type())
			assert.Equal(t, int64(123), video.Size())
			assert.Equal(t, false, video.Processed())
			assert.Equal(t, expectedTime, video.CreatedAt())
		}
	})
}
