package usecases_test

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/matheusvmallmann/plataforma-ead/service-course/application/usecases"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/apptimer"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/ports"
	"github.com/matheusvmallmann/plataforma-ead/service-course/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestProcessVideo(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("When get videos returns error, returns error", func(t *testing.T) {
		// Arrange
		videosRepository := mocks.NewMockVideosRepository(mockCtrl)
		filesService := mocks.NewMockFilesService(mockCtrl)
		videosRepository.EXPECT().Get(ports.GetFilters{Status: "pending"}).Return(nil, errors.New("Test"))
		useCase := usecases.NewProcessVideo(videosRepository, filesService)

		// Act
		err := useCase.Execute()

		// Assert
		assert.Error(t, err)
	})

	t.Run("When get videos returns no one, skip processing", func(t *testing.T) {
		// Arrange
		videosRepository := mocks.NewMockVideosRepository(mockCtrl)
		filesService := mocks.NewMockFilesService(mockCtrl)
		videosRepository.EXPECT().
			Get(ports.GetFilters{Status: "pending"}).
			Return([]*entities.Video{}, nil)
		useCase := usecases.NewProcessVideo(videosRepository, filesService)

		// Act
		err := useCase.Execute()

		// Assert
		assert.Nil(t, err)
	})

	t.Run("When find file resolution return error, set video status to error", func(t *testing.T) {
		// Arrange
		expectedTime := time.Date(2023, time.October, 6, 12, 0, 0, 0, time.UTC)
		fakeTimer := apptimer.NewFakeTimer(expectedTime)
		video, _ := entities.NewVideo(fakeTimer, "123", "/tmp/123.mp4", "mp4", 10000)
		videos := []*entities.Video{video}

		videosRepository := mocks.NewMockVideosRepository(mockCtrl)
		filesService := mocks.NewMockFilesService(mockCtrl)

		videosRepository.EXPECT().Get(ports.GetFilters{Status: "pending"}).Return(videos, nil)
		videosRepository.EXPECT().Update(video).Return(nil)

		filesService.EXPECT().GetResolution("/tmp/123.mp4").Return("", errors.New("test"))

		useCase := usecases.NewProcessVideo(videosRepository, filesService)

		// Act
		err := useCase.Execute()

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, "error", video.Status())
	})

	t.Run("When resolution is invalid, set video status to error", func(t *testing.T) {
		// Arrange
		expectedTime := time.Date(2023, time.October, 6, 12, 0, 0, 0, time.UTC)
		fakeTimer := apptimer.NewFakeTimer(expectedTime)
		video, _ := entities.NewVideo(fakeTimer, "123", "/tmp/123.mp4", "mp4", 10000)
		videos := []*entities.Video{video}

		videosRepository := mocks.NewMockVideosRepository(mockCtrl)
		filesService := mocks.NewMockFilesService(mockCtrl)

		videosRepository.EXPECT().Get(ports.GetFilters{Status: "pending"}).Return(videos, nil)
		videosRepository.EXPECT().Update(video).Return(nil)

		filesService.EXPECT().GetResolution("/tmp/123.mp4").Return("1080xfad", nil)

		useCase := usecases.NewProcessVideo(videosRepository, filesService)

		// Act
		err := useCase.Execute()

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, "error", video.Status())
	})

	t.Run("When all process video returns error, set video status to error", func(t *testing.T) {
		// Arrange
		expectedTime := time.Date(2023, time.October, 6, 12, 0, 0, 0, time.UTC)
		fakeTimer := apptimer.NewFakeTimer(expectedTime)
		video, _ := entities.NewVideo(fakeTimer, "123", "/tmp/123.mp4", "mp4", 10000)
		videos := []*entities.Video{video}

		videosRepository := mocks.NewMockVideosRepository(mockCtrl)
		filesService := mocks.NewMockFilesService(mockCtrl)

		videosRepository.EXPECT().Get(ports.GetFilters{Status: "pending"}).Return(videos, nil)
		videosRepository.EXPECT().Update(video).Return(nil)

		filesService.EXPECT().GetResolution("/tmp/123.mp4").Return("1920x1080", nil)
		filesService.EXPECT().ProcessVideo("/tmp/123.mp4", "/app/tmp/123/480", "640:480").Return(errors.New("test"))
		filesService.EXPECT().ProcessVideo("/tmp/123.mp4", "/app/tmp/123/1080", "1920:1080").Return(errors.New("test"))
		filesService.EXPECT().ProcessVideo("/tmp/123.mp4", "/app/tmp/123/720", "1080:720").Return(errors.New("test"))

		useCase := usecases.NewProcessVideo(videosRepository, filesService)

		// Act
		err := useCase.Execute()

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, "error", video.Status())
	})

	t.Run("When many process video returs error, set video status to success", func(t *testing.T) {
		expectedTime := time.Date(2023, time.October, 6, 12, 0, 0, 0, time.UTC)
		fakeTimer := apptimer.NewFakeTimer(expectedTime)
		video, _ := entities.NewVideo(fakeTimer, "123", "/tmp/123.mp4", "mp4", 10000)
		videos := []*entities.Video{video}

		videosRepository := mocks.NewMockVideosRepository(mockCtrl)
		filesService := mocks.NewMockFilesService(mockCtrl)

		videosRepository.EXPECT().Get(ports.GetFilters{Status: "pending"}).Return(videos, nil)
		videosRepository.EXPECT().Update(video).Return(nil)

		filesService.EXPECT().GetResolution("/tmp/123.mp4").Return("1920x1080", nil)
		filesService.EXPECT().ProcessVideo("/tmp/123.mp4", "/app/tmp/123/480", "640:480").Return(nil)
		filesService.EXPECT().ProcessVideo("/tmp/123.mp4", "/app/tmp/123/1080", "1920:1080").Return(errors.New("test"))
		filesService.EXPECT().ProcessVideo("/tmp/123.mp4", "/app/tmp/123/720", "1080:720").Return(nil)

		useCase := usecases.NewProcessVideo(videosRepository, filesService)

		// Act
		err := useCase.Execute()

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, "success", video.Status())

		if assert.Equal(t, len(video.GetResolutions()), 2) {
			resolution1 := video.GetResolutions()[0]
			assert.Equal(t, "/app/tmp/123/480", resolution1.URL)
			assert.Equal(t, 480, resolution1.Resolution)

			resolution3 := video.GetResolutions()[1]
			assert.Equal(t, "/app/tmp/123/720", resolution3.URL)
			assert.Equal(t, 720, resolution3.Resolution)
		}
	})

	t.Run("When no one process video returs error, set video status to success", func(t *testing.T) {
		expectedTime := time.Date(2023, time.October, 6, 12, 0, 0, 0, time.UTC)
		fakeTimer := apptimer.NewFakeTimer(expectedTime)
		video, _ := entities.NewVideo(fakeTimer, "123", "/tmp/123.mp4", "mp4", 10000)
		videos := []*entities.Video{video}

		videosRepository := mocks.NewMockVideosRepository(mockCtrl)
		filesService := mocks.NewMockFilesService(mockCtrl)

		videosRepository.EXPECT().Get(ports.GetFilters{Status: "pending"}).Return(videos, nil)
		videosRepository.EXPECT().Update(video).Return(nil)

		filesService.EXPECT().GetResolution("/tmp/123.mp4").Return("1920x1080", nil)
		filesService.EXPECT().ProcessVideo("/tmp/123.mp4", "/app/tmp/123/480", "640:480").Return(nil)
		filesService.EXPECT().ProcessVideo("/tmp/123.mp4", "/app/tmp/123/1080", "1920:1080").Return(nil)
		filesService.EXPECT().ProcessVideo("/tmp/123.mp4", "/app/tmp/123/720", "1080:720").Return(nil)

		useCase := usecases.NewProcessVideo(videosRepository, filesService)

		// Act
		err := useCase.Execute()

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, "success", video.Status())

		if assert.Equal(t, len(video.GetResolutions()), 3) {
			resolution1 := video.GetResolutions()[0]
			assert.Equal(t, "/app/tmp/123/480", resolution1.URL)
			assert.Equal(t, 480, resolution1.Resolution)

			resolution3 := video.GetResolutions()[1]
			assert.Equal(t, "/app/tmp/123/720", resolution3.URL)
			assert.Equal(t, 720, resolution3.Resolution)

			resolution2 := video.GetResolutions()[2]
			assert.Equal(t, "/app/tmp/123/1080", resolution2.URL)
			assert.Equal(t, 1080, resolution2.Resolution)
		}
	})
}
