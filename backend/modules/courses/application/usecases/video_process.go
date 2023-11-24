package usecases

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"

	errs "github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/errors"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/utils"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/ports"
)

type ProcessVideo struct {
	videosRepository ports.VideosRepository
	filesService     ports.FilesService
}

type ProcessVideoInput struct {
}

type resolutionType struct {
	resolution string
	width      int32
	height     int32
	folderName string
	bandWidth  int64
}

type ProcessResolutionChannel struct {
	URL        string
	resolution resolutionType
}

var resolutions = []resolutionType{
	{resolution: "1080", width: 1920, height: 1080, folderName: "1080", bandWidth: 5000000},
	{resolution: "720", width: 1080, height: 720, folderName: "720", bandWidth: 2800000},
	{resolution: "480", width: 640, height: 480, folderName: "480", bandWidth: 1400000},
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
	resultChannel := make(chan ProcessResolutionChannel, threads)

	for _, resolution := range resolutions {
		wg.Add(1)
		go p.processResolution(&wg, video, resolution, resultChannel)
	}

	wg.Wait()
	close(resultChannel)

	var resolutionsURL []string
	var successResolutions []resolutionType
	for resolution := range resultChannel {
		resolutionsURL = append(resolutionsURL, resolution.URL)
		successResolutions = append(successResolutions, resolution.resolution)
	}

	if len(resolutionsURL) < 1 {
		return video.SetStatus("error")
	}

	url, err := p.createQualityFile(video, successResolutions)

	if err != nil {
		log.Printf("Error to create quality file: %s", err)
		return video.SetStatus("error")
	}

	for index, res := range successResolutions {
		resolution := entities.VideoResolution{
			URL:        resolutionsURL[index],
			Resolution: res.resolution,
		}
		video.AddResolution(resolution)
	}

	p.filesService.Delete(video.TmpUrl())

	return video.SetUrl(url).SetStatus("success")
}

func (p *ProcessVideo) processResolution(
	wg *sync.WaitGroup,
	video *entities.Video,
	videoResolution resolutionType,
	resultChannel chan ProcessResolutionChannel,
) error {
	defer wg.Done()
	resolution := fmt.Sprintf("%d:%d", videoResolution.width, videoResolution.height)
	log.Printf("[%s-%s] Starting process resolution...", video.Id(), resolution)
	folderPath := fmt.Sprintf("/videos/%s/%s", video.Id(), videoResolution.folderName)

	if err := p.filesService.ProcessVideo(video.TmpUrl(), folderPath, resolution); err != nil {
		log.Printf("[%s-%s] Error to process video: %s", video.Id(), resolution, err)
		return err
	}

	resultChannel <- ProcessResolutionChannel{
		URL:        folderPath,
		resolution: videoResolution,
	}
	log.Printf("[%s-%s] Video processed successfully.", video.Id(), resolution)
	return nil
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

func (p *ProcessVideo) createQualityFile(video *entities.Video, resolutions []resolutionType) (string, error) {
	sorteredResolutions := utils.SortSlice[resolutionType](resolutions, func(i, j int) bool {
		return resolutions[i].height < resolutions[j].height
	})

	fileText := "#EXTM3U\n#EXT-X-VERSION:3\n"

	for _, res := range sorteredResolutions {
		resolution := fmt.Sprintf("%dx%d", res.width, res.height)
		line := fmt.Sprintf("#EXT-X-STREAM-INF:BANDWIDTH=%d,RESOLUTION=%s\n%s/index.m3u8\n", res.bandWidth, resolution, res.folderName)
		fileText += line
	}

	videoUrl := fmt.Sprintf("/videos/%s/playlist.m3u8", video.Id())

	file, err := p.filesService.CreateFile(ports.FileInfo{
		Url:  videoUrl,
		Type: "m3u8",
	})
	if err != nil {
		return "", err
	}

	if err := file.WriteString(fileText); err != nil {
		return "", err
	}

	file.Close()

	return videoUrl, nil
}
