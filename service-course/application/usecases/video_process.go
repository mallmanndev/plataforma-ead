package usecases

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"

	errs "github.com/matheusvmallmann/plataforma-ead/service-course/application/errors"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/ports"
)

type ProcessVideo struct {
	videosRepository ports.VideosRepository
	filesService     ports.FilesService
}

type ProcessVideoInput struct {
}

type resolutionType struct {
	resolution         string
	completeResolution string
	folderName         string
}

func NewProcessVideo(
	videosRepository ports.VideosRepository,
	filesService ports.FilesService,
) *ProcessVideo {
	return &ProcessVideo{
		videosRepository: videosRepository,
		filesService:     filesService,
	}
}

func (p *ProcessVideo) Execute() error {

	videos, err := p.videosRepository.Get(ports.GetFilters{Status: "pending"})
	if err != nil {
		log.Println("Error getting videos to process: ", err)
		return err
	}
	if len(videos) < 1 {
		return nil
	}

	for _, video := range videos {
		processVideo := p.processVideo(video)
		if err := p.videosRepository.Update(processVideo); err != nil {
			log.Println("Error to update video: ", err)
			return errs.NewProcessVideoUseCaseError("Could not update video", err)
		}
		log.Printf("[%s] Video updated successfully.", video.Id())
	}
	return nil
}

func (p *ProcessVideo) processVideo(video *entities.Video) *entities.Video {
	log.Printf("[%s] Starting process video...", video.Id())

	videoResolution, err := p.filesService.GetResolution(video.TmpUrl())
	if err != nil {
		log.Printf("Error to process videos: %s", err)
		return video.SetStatus("error")
	}

	resolutionPixels := strings.Split(videoResolution, "x")
	heightResolution := resolutionPixels[1]

	resolutions, err := p.filterResolutions(heightResolution)
	if err != nil {
		log.Printf("Error to process videos: %s", err)
		return video.SetStatus("error")
	}

	threads := len(resolutions)
	var wg sync.WaitGroup
	resultChannel := make(chan entities.VideoResolution, threads)

	for _, resolution := range resolutions {
		wg.Add(1)
		go p.processResolution(&wg, video, resolution, resultChannel)
	}

	wg.Wait()
	close(resultChannel)

	video.SetStatus("error")
	for resolution := range resultChannel {
		video.AddResolution(resolution)
		video.SetStatus("success")
	}

	return video
}

func (p *ProcessVideo) processResolution(
	wg *sync.WaitGroup,
	video *entities.Video,
	videoResolution resolutionType,
	resultChannel chan entities.VideoResolution,
) error {
	defer wg.Done()
	log.Printf("[%s-%s] Starting process resolution...", video.Id(), videoResolution.completeResolution)
	folderPath := fmt.Sprintf("/videos/%s/%s", video.Id(), videoResolution.folderName)

	if err := p.filesService.ProcessVideo(video.TmpUrl(), folderPath, videoResolution.completeResolution); err != nil {
		log.Printf("[%s-%s] Error to process video: %s", video.Id(), videoResolution.completeResolution, err)
		return err
	}

	resultChannel <- entities.VideoResolution{
		URL:        folderPath,
		Resolution: videoResolution.resolution,
	}
	log.Printf("[%s-%s] Video processed successfully.", video.Id(), videoResolution.completeResolution)
	return nil
}

var resolutions = []resolutionType{
	{resolution: "1080", completeResolution: "1920:1080", folderName: "1080"},
	{resolution: "720", completeResolution: "1080:720", folderName: "720"},
	{resolution: "480", completeResolution: "640:480", folderName: "480"},
}

func (p *ProcessVideo) filterResolutions(fileResolution string) ([]resolutionType, error) {
	fileResolutionInt, _ := strconv.Atoi(fileResolution)
	filterResolution := []resolutionType{}
	hasResolution := false

	for _, resolution := range resolutions {
		resolutionInt, _ := strconv.Atoi(resolution.resolution)
		if resolutionInt == fileResolutionInt {
			hasResolution = true
		}

		if fileResolutionInt < resolutionInt {
			continue
		}
		filterResolution = append(filterResolution, resolution)
	}

	if !hasResolution {
		return nil, errs.NewProcessVideoUseCaseError("Invalid file resolution.", nil)
	}

	return filterResolution, nil
}
