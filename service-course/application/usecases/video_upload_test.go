package usecases_test

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/matheusvmallmann/plataforma-ead/service-course/application/usecases"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/apptimer"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/ports"
	"github.com/matheusvmallmann/plataforma-ead/service-course/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestVideoUpload_CreateFile(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	filesService := mocks.NewMockFilesService(mockCtrl)
	filesServiceTwo := mocks.NewMockFilesService(mockCtrl)
	uuidService := mocks.NewMockUUIDService(mockCtrl)
	videosRepository := mocks.NewMockVideosRepository(mockCtrl)

	t.Run("Should return error when type is invalid", func(t *testing.T) {
		// GIVEN
		expectedTime := time.Date(2023, time.October, 6, 12, 0, 0, 0, time.UTC)
		fakeTimer := apptimer.NewFakeTimer(expectedTime)
		useCase := usecases.NewVideoUpload(filesService, uuidService, videosRepository, fakeTimer)

		uuidService.EXPECT().Generate().Return("123456789")

		// WHEN
		_, err := useCase.CreateFile("mp", 10000)

		// THEN
		assert.ErrorContains(t, err, "[Video] Invalid 'type': must be mp4.")
	})

	t.Run("Should return error when fileService returns error", func(t *testing.T) {
		// GIVEN
		expectedTime := time.Date(2023, time.October, 6, 12, 0, 0, 0, time.UTC)
		fakeTimer := apptimer.NewFakeTimer(expectedTime)
		useCase := usecases.NewVideoUpload(filesService, uuidService, videosRepository, fakeTimer)

		uuidService.EXPECT().Generate().Return("123456789")
		filesService.EXPECT().
			CreateFile(ports.FileInfo{Url: "/videos/tmp/123456789.mp4", Type: "mp4", Size: 10000}).
			Return(nil, errors.New("Test!"))

		// WHEN
		_, err := useCase.CreateFile("mp4", 10000)

		// THEN
		t.Log(err)
		assert.ErrorContains(t, err, "[Video Upload] Could not create video: Test!")
	})

	t.Run("Should create a new video file", func(t *testing.T) {
		// GIVEN
		expectedTime := time.Date(2023, time.October, 6, 12, 0, 0, 0, time.UTC)
		fakeTimer := apptimer.NewFakeTimer(expectedTime)
		useCase := usecases.NewVideoUpload(filesService, uuidService, videosRepository, fakeTimer)

		uuidService.EXPECT().Generate().Return("123456789")
		filesService.EXPECT().
			CreateFile(ports.FileInfo{Url: "/videos/tmp/123456789.mp4", Type: "mp4", Size: 10000}).
			Return(filesServiceTwo, nil)

		// WHEN
		file, err := useCase.CreateFile("mp4", 10000)

		// THEN
		assert.Nil(t, err)
		if assert.NotNil(t, file) {
			assert.Equal(t, "123456789", file.Video().Id())
			assert.Equal(t, "/videos/tmp/123456789.mp4", file.Video().TmpUrl())
			assert.Equal(t, "mp4", file.Video().Type())
			assert.Equal(t, "pending", file.Video().Status())
		}
	})
}

func TestVideoUpload_SendChunk(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	filesService := mocks.NewMockFilesService(mockCtrl)
	filesServiceTwo := mocks.NewMockFilesService(mockCtrl)
	uuidService := mocks.NewMockUUIDService(mockCtrl)
	videosRepository := mocks.NewMockVideosRepository(mockCtrl)

	t.Run("Should return error when not create a file", func(t *testing.T) {
		//GIVEN
		expectedTime := time.Date(2023, time.October, 6, 12, 0, 0, 0, time.UTC)
		fakeTimer := apptimer.NewFakeTimer(expectedTime)
		useCase := usecases.NewVideoUpload(filesService, uuidService, videosRepository, fakeTimer)

		// WHEN
		err := useCase.SendChunk([]byte("test"))

		// THEN
		assert.ErrorContains(t, err, "[Video Upload] Video has not been created.")
	})

	t.Run("Should return error when fileService returns error", func(t *testing.T) {
		// GIVEN
		expectedTime := time.Date(2023, time.October, 6, 12, 0, 0, 0, time.UTC)
		fakeTimer := apptimer.NewFakeTimer(expectedTime)
		useCase := usecases.NewVideoUpload(filesService, uuidService, videosRepository, fakeTimer)

		uuidService.EXPECT().Generate().Return("123456789")
		filesService.EXPECT().
			CreateFile(ports.FileInfo{Url: "/videos/tmp/123456789.mp4", Type: "mp4", Size: 10000}).
			Return(filesServiceTwo, nil)
		filesServiceTwo.EXPECT().SendChunk([]byte("test")).Return(errors.New("Test!"))

		// WHEN
		file, _ := useCase.CreateFile("mp4", 10000)

		// THEN
		err := file.SendChunk([]byte("test"))
		assert.ErrorContains(t, err, "[Video Upload] Could not send video chunk: Test!")
	})
}

func TestVideoUpload_Execute(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	filesService := mocks.NewMockFilesService(mockCtrl)
	filesServiceTwo := mocks.NewMockFilesService(mockCtrl)
	uuidService := mocks.NewMockUUIDService(mockCtrl)
	videosRepository := mocks.NewMockVideosRepository(mockCtrl)

	t.Run("Should return error when not create a file", func(t *testing.T) {
		// GIVEN
		expectedTime := time.Date(2023, time.October, 6, 12, 0, 0, 0, time.UTC)
		fakeTimer := apptimer.NewFakeTimer(expectedTime)
		useCase := usecases.NewVideoUpload(filesService, uuidService, videosRepository, fakeTimer)

		// THEN
		_, err := useCase.Execute()

		// WHEN
		assert.ErrorContains(t, err, "[Video Upload] Video has not been created.")
	})

	t.Run("Should return error when file size is different", func(t *testing.T) {
		// GIVEN
		expectedTime := time.Date(2023, time.October, 6, 12, 0, 0, 0, time.UTC)
		fakeTimer := apptimer.NewFakeTimer(expectedTime)
		useCase := usecases.NewVideoUpload(filesService, uuidService, videosRepository, fakeTimer)

		chunk := []byte("test")

		uuidService.EXPECT().Generate().Return("123456789")
		filesService.EXPECT().
			CreateFile(ports.FileInfo{Url: "/videos/tmp/123456789.mp4", Type: "mp4", Size: 25}).
			Return(filesServiceTwo, nil)
		filesServiceTwo.EXPECT().SendChunk(chunk).Return(nil)
		filesServiceTwo.EXPECT().SendChunk(chunk).Return(nil)
		filesServiceTwo.EXPECT().Close().Return(nil)
		filesServiceTwo.EXPECT().Remove().Return(nil)

		// WHEN
		file, _ := useCase.CreateFile("mp4", 25)
		file.SendChunk(chunk)
		file.SendChunk(chunk)

		// THEN
		_, err := file.Execute()
		assert.ErrorContains(t, err, "[Video Upload] Invalid 'size': Expected 25 bytes, Received 8 bytes.")
	})

	t.Run("When close file returns error, return error", func(t *testing.T) {
		// GIVEN
		expectedTime := time.Date(2023, time.October, 6, 12, 0, 0, 0, time.UTC)
		fakeTimer := apptimer.NewFakeTimer(expectedTime)
		useCase := usecases.NewVideoUpload(filesService, uuidService, videosRepository, fakeTimer)

		chunk := []byte("test")

		uuidService.EXPECT().Generate().Return("123456789")
		filesService.EXPECT().
			CreateFile(ports.FileInfo{Url: "/videos/tmp/123456789.mp4", Type: "mp4", Size: 25}).
			Return(filesServiceTwo, nil)
		filesServiceTwo.EXPECT().SendChunk(chunk).Return(nil)
		filesServiceTwo.EXPECT().SendChunk(chunk).Return(nil)
		filesServiceTwo.EXPECT().Close().Return(errors.New("test"))

		// WHEN
		file, _ := useCase.CreateFile("mp4", 25)
		file.SendChunk(chunk)
		file.SendChunk(chunk)

		// THEN
		_, err := file.Execute()
		assert.ErrorContains(t, err, "[Video Upload] Could not save video: test")
	})

	t.Run("Should return error when not insert video in repository", func(t *testing.T) {
		// GIVEN
		expectedTime := time.Date(2023, time.October, 6, 12, 0, 0, 0, time.UTC)
		fakeTimer := apptimer.NewFakeTimer(expectedTime)
		useCase := usecases.NewVideoUpload(filesService, uuidService, videosRepository, fakeTimer)

		chunk := []byte("teste")
		uuidService.EXPECT().Generate().Return("123456789")

		filesService.EXPECT().
			CreateFile(ports.FileInfo{Url: "/videos/tmp/123456789.mp4", Type: "mp4", Size: 25}).
			Return(filesServiceTwo, nil)
		filesServiceTwo.EXPECT().SendChunk(chunk).Return(nil)
		filesServiceTwo.EXPECT().SendChunk(chunk).Return(nil)
		filesServiceTwo.EXPECT().SendChunk(chunk).Return(nil)
		filesServiceTwo.EXPECT().SendChunk(chunk).Return(nil)
		filesServiceTwo.EXPECT().SendChunk(chunk).Return(nil)
		filesServiceTwo.EXPECT().Close().Return(nil)
		filesServiceTwo.EXPECT().Remove().Return(nil)
		videosRepository.EXPECT().Create(gomock.Any()).Return(errors.New("Test!"))

		// WHEN
		file, _ := useCase.CreateFile("mp4", 25)
		file.SendChunk(chunk)
		file.SendChunk(chunk)
		file.SendChunk(chunk)
		file.SendChunk(chunk)
		file.SendChunk(chunk)
		_, err := file.Execute()

		// THEN
		assert.ErrorContains(t, err, "[Video Upload] Could not save video: Test!")
	})
}
