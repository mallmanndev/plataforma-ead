package usecases

import (
	"fmt"

	errs "github.com/matheusvmallmann/plataforma-ead/service-course/application/errors"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/apptimer"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/ports"
)

type VideoUpload struct {
	filesService ports.FilesService
	uuidService  ports.UUIDService
	repository   ports.VideosRepository
	video        *entities.Video
	size         int64
	chunkSize    int64
	timer        apptimer.Timer
}

func NewVideoUpload(
	FilesService ports.FilesService,
	UUIDService ports.UUIDService,
	VideosRepository ports.VideosRepository,
	Timer apptimer.Timer,
) *VideoUpload {
	return &VideoUpload{
		filesService: FilesService,
		uuidService:  UUIDService,
		repository:   VideosRepository,
		timer:        Timer,
	}
}

func (v *VideoUpload) CreateFile(Type string, Size int64, UserId string) (*VideoUpload, error) {
	videoId := v.uuidService.Generate()
	videoUrl := fmt.Sprintf("/videos/tmp/%s.%s", videoId, Type)
	video, err := entities.NewVideo(v.timer, videoId, videoUrl, Type, Size, UserId)
	if err != nil {
		return nil, err
	}

	newService, err := v.filesService.CreateFile(ports.FileInfo{
		Url:  videoUrl,
		Size: Size,
		Type: Type,
	})
	if err != nil {
		return nil, errs.NewVideoUploadUseCaseError("Could not create video", err)
	}

	newVideoUpload := &VideoUpload{
		video:        video,
		filesService: newService,
		repository:   v.repository,
	}
	return newVideoUpload, nil
}

func (v *VideoUpload) SendChunk(data []byte) error {
	if v.video == nil {
		return errs.NewVideoUploadUseCaseError("Video has not been created", nil)
	}

	if err := v.filesService.SendChunk(data); err != nil {
		return errs.NewVideoUploadUseCaseError("Could not send video chunk", err)
	}
	v.size += int64(len(data))
	v.chunkSize = int64(len(data))

	return nil
}

func (v *VideoUpload) Execute() (*entities.Video, error) {
	if v.video == nil {
		return nil, errs.NewVideoUploadUseCaseError("Video has not been created", nil)
	}

	if err := v.filesService.Close(); err != nil {
		return nil, errs.NewVideoUploadUseCaseError("Could not save video", err)
	}

	sizeDiff := v.video.Size() - v.size
	if v.size == 0 || sizeDiff > v.chunkSize {
		v.filesService.Remove()
		return nil, errs.NewInvalidAttributeError(
			"Video Upload",
			"size",
			fmt.Sprintf("Expected %d bytes, Received %d bytes", v.video.Size(), v.size),
		)
	}

	if err := v.repository.Create(v.video); err != nil {
		v.filesService.Remove()
		return nil, errs.NewVideoUploadUseCaseError("Could not save video", err)
	}

	return v.video, nil
}

func (v *VideoUpload) Video() *entities.Video {
	return v.video
}
