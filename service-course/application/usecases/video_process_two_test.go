package usecases_test

import (
	"errors"
	"fmt"
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

func TestProcessVideoTwo(t *testing.T) {
	var videosFolder = "/videos"

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("When get videos returns error, returns error", func(t *testing.T) {
		// Arrange
		videosRepository := mocks.NewMockVideosRepository(mockCtrl)
		filesService := mocks.NewMockFilesService(mockCtrl)
		videosRepository.EXPECT().Get(ports.GetFilters{Status: "pending"}).Return(nil, errors.New("Test"))
		useCase := usecases.NewProcessVideoTwo(videosRepository, filesService)

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
		useCase := usecases.NewProcessVideoTwo(videosRepository, filesService)

		// Act
		err := useCase.Execute()

		// Assert
		assert.Nil(t, err)
	})

	t.Run("When find file resolution return error, set video status to error", func(t *testing.T) {
		// Arrange
		expectedTime := time.Date(2023, time.October, 6, 12, 0, 0, 0, time.UTC)
		fakeTimer := apptimer.NewFakeTimer(expectedTime)
		video, _ := entities.NewVideo(fakeTimer, "123", "/tmp/123.mp4", "mp4", 10000, "user_id")
		videos := []*entities.Video{video}

		videosRepository := mocks.NewMockVideosRepository(mockCtrl)
		filesService := mocks.NewMockFilesService(mockCtrl)

		videosRepository.EXPECT().Get(ports.GetFilters{Status: "pending"}).Return(videos, nil)
		videosRepository.EXPECT().Update(video).Return(nil)

		filesService.EXPECT().GetResolution("/tmp/123.mp4").Return("", errors.New("test"))

		useCase := usecases.NewProcessVideoTwo(videosRepository, filesService)

		// Act
		err := useCase.Execute()

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, "error", video.Status())
	})

	t.Run("When_resolution_is_invalid_set_video_status_to_error", func(t *testing.T) {
		// Arrange
		expectedTime := time.Date(2023, time.October, 6, 12, 0, 0, 0, time.UTC)
		fakeTimer := apptimer.NewFakeTimer(expectedTime)
		video, _ := entities.NewVideo(fakeTimer, "123", "/tmp/123.mp4", "mp4", 10000, "user_id")
		videos := []*entities.Video{video}

		videosRepository := mocks.NewMockVideosRepository(mockCtrl)
		filesService := mocks.NewMockFilesService(mockCtrl)
		filesServiceWithFile := mocks.NewMockFilesService(mockCtrl)

		filesService.EXPECT().
			CreateFile(ports.FileInfo{Url: "/videos/123/playlist.m3u8", Type: "m3u8", Size: 0}).
			Return(filesServiceWithFile, nil)
		filesServiceWithFile.EXPECT().WriteString("#EXTM3U\n#EXT-X-VERSION:3\n").Return(nil)

		videosRepository.EXPECT().Get(ports.GetFilters{Status: "pending"}).Return(videos, nil)
		videosRepository.EXPECT().Update(video).Return(nil)

		filesService.EXPECT().GetResolution("/tmp/123.mp4").Return("1080xfad", nil)

		useCase := usecases.NewProcessVideoTwo(videosRepository, filesService)

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
		video, _ := entities.NewVideo(fakeTimer, "123", "/tmp/123.mp4", "mp4", 10000, "user_id")
		videos := []*entities.Video{video}

		videosRepository := mocks.NewMockVideosRepository(mockCtrl)
		filesService := mocks.NewMockFilesService(mockCtrl)
		filesServiceWithFile := mocks.NewMockFilesService(mockCtrl)

		videosRepository.EXPECT().Get(ports.GetFilters{Status: "pending"}).Return(videos, nil)
		videosRepository.EXPECT().Update(video).Return(nil)
		videosRepository.EXPECT().Update(video).Return(nil)
		videosRepository.EXPECT().Update(video).Return(nil)
		videosRepository.EXPECT().Update(video).Return(nil)

		filesService.EXPECT().
			CreateFile(ports.FileInfo{Url: "/videos/123/playlist.m3u8", Type: "m3u8", Size: 0}).
			Return(filesServiceWithFile, nil)
		filesServiceWithFile.EXPECT().WriteString("#EXTM3U\n#EXT-X-VERSION:3\n").Return(nil)
		filesServiceWithFile.EXPECT().Close()

		filesService.EXPECT().GetResolution("/tmp/123.mp4").Return("1920x1080", nil)
		filesService.EXPECT().ProcessVideo("/tmp/123.mp4", videosFolder+"/123/480", "640:480").Return(errors.New("test"))
		filesService.EXPECT().ProcessVideo("/tmp/123.mp4", videosFolder+"/123/1080", "1920:1080").Return(errors.New("test"))
		filesService.EXPECT().ProcessVideo("/tmp/123.mp4", videosFolder+"/123/720", "1080:720").Return(errors.New("test"))

		useCase := usecases.NewProcessVideoTwo(videosRepository, filesService)

		// Act
		err := useCase.Execute()

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, "error", video.Status())
	})

	t.Run("When many process video returs error, set video status to success", func(t *testing.T) {
		expectedTime := time.Date(2023, time.October, 6, 12, 0, 0, 0, time.UTC)
		fakeTimer := apptimer.NewFakeTimer(expectedTime)
		video, _ := entities.NewVideo(fakeTimer, "123", "/tmp/123.mp4", "mp4", 10000, "user_id")
		videos := []*entities.Video{video}

		videosRepository := mocks.NewMockVideosRepository(mockCtrl)
		filesService := mocks.NewMockFilesService(mockCtrl)
		filesServiceWithFile := mocks.NewMockFilesService(mockCtrl)

		videosRepository.EXPECT().Get(ports.GetFilters{Status: "pending"}).Return(videos, nil)
		videosRepository.EXPECT().Update(video).Return(nil)
		videosRepository.EXPECT().Update(video).Return(nil)
		videosRepository.EXPECT().Update(video).Return(nil)
		videosRepository.EXPECT().Update(video).Return(nil)

		filesService.EXPECT().GetResolution("/tmp/123.mp4").Return("1920x1080", nil)
		filesService.EXPECT().ProcessVideo("/tmp/123.mp4", videosFolder+"/123/480", "640:480").Return(nil)
		filesService.EXPECT().ProcessVideo("/tmp/123.mp4", videosFolder+"/123/1080", "1920:1080").Return(errors.New("test"))
		filesService.EXPECT().ProcessVideo("/tmp/123.mp4", videosFolder+"/123/720", "1080:720").Return(nil)

		filesServiceWithFile.
			EXPECT().
			WriteString("#EXTM3U\n#EXT-X-VERSION:3\n")
		filesServiceWithFile.
			EXPECT().
			WriteString(fmt.Sprintf("#EXT-X-STREAM-INF:BANDWIDTH=%d,RESOLUTION=%s\n%s/index.m3u8\n", 1400000, "640x480", "480"))
		filesServiceWithFile.
			EXPECT().
			WriteString(fmt.Sprintf("#EXT-X-STREAM-INF:BANDWIDTH=%d,RESOLUTION=%s\n%s/index.m3u8\n", 2800000, "1080x720", "720"))

		filesService.EXPECT().
			CreateFile(ports.FileInfo{Url: "/videos/123/playlist.m3u8", Type: "m3u8", Size: 0}).
			Return(filesServiceWithFile, nil)

		filesServiceWithFile.EXPECT().Close()
		filesService.EXPECT().Delete("/tmp/123.mp4").Return(nil)

		useCase := usecases.NewProcessVideoTwo(videosRepository, filesService)

		// Act
		err := useCase.Execute()

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, "success", video.Status())

		if assert.Equal(t, len(video.GetResolutions()), 3) {
			resolution1 := video.GetResolutions()[0]
			assert.Equal(t, videosFolder+"/123/480", resolution1.URL)
			assert.Equal(t, "480", resolution1.Resolution)
			assert.Equal(t, "success", resolution1.Status)

			resolution2 := video.GetResolutions()[1]
			assert.Equal(t, videosFolder+"/123/720", resolution2.URL)
			assert.Equal(t, "720", resolution2.Resolution)
			assert.Equal(t, "success", resolution2.Status)

			resolution3 := video.GetResolutions()[2]
			assert.Equal(t, "", resolution3.URL)
			assert.Equal(t, "1080", resolution3.Resolution)
			assert.Equal(t, "error", resolution3.Status)
		}
	})

	t.Run("When no one process video returs error, set video status to success", func(t *testing.T) {
		expectedTime := time.Date(2023, time.October, 6, 12, 0, 0, 0, time.UTC)
		fakeTimer := apptimer.NewFakeTimer(expectedTime)
		video, _ := entities.NewVideo(fakeTimer, "123", "/tmp/123.mp4", "mp4", 10000, "user_id")
		videos := []*entities.Video{video}

		videosRepository := mocks.NewMockVideosRepository(mockCtrl)
		filesService := mocks.NewMockFilesService(mockCtrl)
		filesServiceWithFile := mocks.NewMockFilesService(mockCtrl)

		videosRepository.EXPECT().Get(ports.GetFilters{Status: "pending"}).Return(videos, nil)
		videosRepository.EXPECT().Update(video).Return(nil)
		videosRepository.EXPECT().Update(video).Return(nil)
		videosRepository.EXPECT().Update(video).Return(nil)
		videosRepository.EXPECT().Update(video).Return(nil)

		filesService.EXPECT().GetResolution("/tmp/123.mp4").Return("1920x1080", nil)
		filesService.EXPECT().ProcessVideo("/tmp/123.mp4", videosFolder+"/123/480", "640:480").Return(nil)
		filesService.EXPECT().ProcessVideo("/tmp/123.mp4", videosFolder+"/123/720", "1080:720").Return(nil)
		filesService.EXPECT().ProcessVideo("/tmp/123.mp4", videosFolder+"/123/1080", "1920:1080").Return(nil)

		filesServiceWithFile.
			EXPECT().
			WriteString("#EXTM3U\n#EXT-X-VERSION:3\n")
		filesServiceWithFile.
			EXPECT().
			WriteString(fmt.Sprintf("#EXT-X-STREAM-INF:BANDWIDTH=%d,RESOLUTION=%s\n%s/index.m3u8\n", 1400000, "640x480", "480"))
		filesServiceWithFile.
			EXPECT().
			WriteString(fmt.Sprintf("#EXT-X-STREAM-INF:BANDWIDTH=%d,RESOLUTION=%s\n%s/index.m3u8\n", 2800000, "1080x720", "720"))
		filesServiceWithFile.
			EXPECT().
			WriteString(fmt.Sprintf("#EXT-X-STREAM-INF:BANDWIDTH=%d,RESOLUTION=%s\n%s/index.m3u8\n", 5000000, "1920x1080", "1080"))

		filesService.EXPECT().
			CreateFile(ports.FileInfo{Url: "/videos/123/playlist.m3u8", Type: "m3u8", Size: 0}).
			Return(filesServiceWithFile, nil)

		filesServiceWithFile.EXPECT().Close()
		filesService.EXPECT().Delete("/tmp/123.mp4").Return(nil)

		useCase := usecases.NewProcessVideoTwo(videosRepository, filesService)

		// Act
		err := useCase.Execute()

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, "success", video.Status())

		if assert.Equal(t, len(video.GetResolutions()), 3) {
			resolution1 := video.GetResolutions()[0]
			assert.Equal(t, videosFolder+"/123/480", resolution1.URL)
			assert.Equal(t, "480", resolution1.Resolution)
			assert.Equal(t, "success", resolution1.Status)

			resolution2 := video.GetResolutions()[1]
			assert.Equal(t, videosFolder+"/123/720", resolution2.URL)
			assert.Equal(t, "720", resolution2.Resolution)
			assert.Equal(t, "success", resolution2.Status)

			resolution3 := video.GetResolutions()[2]
			assert.Equal(t, videosFolder+"/123/1080", resolution3.URL)
			assert.Equal(t, "1080", resolution3.Resolution)
			assert.Equal(t, "success", resolution3.Status)
		}
	})
}
