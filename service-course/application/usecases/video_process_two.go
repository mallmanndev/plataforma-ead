package usecases

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	errs "github.com/matheusvmallmann/plataforma-ead/service-course/application/errors"
	"github.com/matheusvmallmann/plataforma-ead/service-course/application/utils"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/ports"
)

type ProcessVideoTwo struct {
	videosRepository ports.VideosRepository
	filesService     ports.FilesService
}

type ProcessVideoTwoInput struct {
}

type resolutionTypeTwo struct {
	resolution string
	width      int32
	height     int32
	folderName string
	bandWidth  int64
}

var resolutionsTwo = []resolutionTypeTwo{
	{resolution: "1080", width: 1920, height: 1080, folderName: "1080", bandWidth: 5000000},
	{resolution: "720", width: 1080, height: 720, folderName: "720", bandWidth: 2800000},
	{resolution: "480", width: 640, height: 480, folderName: "480", bandWidth: 1400000},
}

func NewProcessVideoTwo(
	videosRepository ports.VideosRepository,
	filesService ports.FilesService,
) *ProcessVideoTwo {
	return &ProcessVideoTwo{
		videosRepository: videosRepository,
		filesService:     filesService,
	}
}

func (p *ProcessVideoTwo) Execute() error {
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

func (p *ProcessVideoTwo) processVideo(video *entities.Video) *entities.Video {
	log.Printf("[%s] Starting process video...", video.Id())

	videoResolution, err := p.filesService.GetResolution(video.TmpUrl())
	if err != nil {
		log.Printf("Error to process videos: %s", err)
		return video.SetStatus("error")
	}

	resolutionPixels := strings.Split(videoResolution, "x")
	heightResolution := resolutionPixels[1]

	url := fmt.Sprintf("/videos/%s/playlist.m3u8", video.Id())
	video.SetUrl(url)
	file, err := p.filesService.CreateFile(
		ports.FileInfo{Url: url, Type: "m3u8"},
	)
	if err != nil {
		log.Printf("Error to process videos: %s", err)
		return video.SetStatus("error")
	}

	file.WriteString("#EXTM3U\n#EXT-X-VERSION:3\n")

	resolutions, err := p.filterResolutions(heightResolution)
	if err != nil {
		log.Printf("Error to process videos: %s", err)
		return video.SetStatus("error")
	}

	sorteredResolutions := utils.SortSlice[resolutionTypeTwo](resolutionsTwo, func(i, j int) bool {
		return resolutions[i].height < resolutions[j].height
	})

	success := false
	for _, resolution := range sorteredResolutions {
		url, err := p.processResolution(video, resolution)
		if err != nil {
			video.AddResolution(
				entities.VideoResolution{
					URL:        url,
					Resolution: resolution.resolution,
					Status:     "error",
				},
			)
			if err := p.videosRepository.Update(video); err != nil {
				log.Println("Error to update video: ", err)
			}
			continue
		}

		res := fmt.Sprintf("%dx%d", resolution.width, resolution.height)
		line := fmt.Sprintf(
			"#EXT-X-STREAM-INF:BANDWIDTH=%d,RESOLUTION=%s\n%s/index.m3u8\n",
			resolution.bandWidth,
			res,
			resolution.folderName,
		)
		file.WriteString(line)

		video.SetStatus("success")
		video.AddResolution(
			entities.VideoResolution{
				URL:        url,
				Resolution: resolution.resolution,
				Status:     "success",
			},
		)

		if err := p.videosRepository.Update(video); err != nil {
			log.Println("Error to update video: ", err)
			continue
		}
		success = true
	}

	file.Close()
	if !success {
		return video.SetStatus("error")
	}

	p.filesService.Delete(video.TmpUrl())
	return video
}

func (p *ProcessVideoTwo) processResolution(
	video *entities.Video,
	videoResolution resolutionTypeTwo,
) (string, error) {
	resolution := fmt.Sprintf("%d:%d", videoResolution.width, videoResolution.height)
	log.Printf("[%s-%s] Starting process resolution...", video.Id(), resolution)
	folderPath := fmt.Sprintf("/videos/%s/%s", video.Id(), videoResolution.folderName)

	if err := p.filesService.ProcessVideo(video.TmpUrl(), folderPath, resolution); err != nil {
		log.Printf("[%s-%s] Error to process video: %s", video.Id(), resolution, err)
		return "", err
	}

	log.Printf("[%s-%s] Video processed successfully.", video.Id(), resolution)
	return folderPath, nil
}

func (p *ProcessVideoTwo) filterResolutions(fileResolution string) ([]resolutionTypeTwo, error) {
	fileResolutionInt, _ := strconv.Atoi(fileResolution)
	filterResolution := []resolutionTypeTwo{}
	hasResolution := false

	for _, resolution := range resolutionsTwo {
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
