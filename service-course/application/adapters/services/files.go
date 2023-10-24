package services

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/ports"
)

type FilesService struct {
	url  string
	size int64
	file *os.File
}

func NewFilesService() ports.FilesService {
	return &FilesService{}
}

func (r *FilesService) CreateFile(File ports.FileInfo) (ports.FilesService, error) {
	file, err := os.Create(File.Url)
	if err != nil {
		return nil, err
	}

	service := &FilesService{
		url:  File.Url,
		size: File.Size,
		file: file,
	}

	return service, nil
}

func (r *FilesService) SendChunk(chunk []byte) error {
	_, err := r.file.Write(chunk)
	return err
}

func (r *FilesService) WriteString(content string) error {
	if _, err := r.file.WriteString(content); err != nil {
		return err
	}
	return nil
}

func (r *FilesService) Close() error {
	return r.file.Close()
}

func (r *FilesService) Remove() error {
	return os.Remove(r.url)
}

func (r *FilesService) GetResolution(Url string) (string, error) {
	cmd := exec.Command(
		"ffprobe",
		"-v", "error",
		"-select_streams", "v:0",
		"-show_entries",
		"stream=width,height",
		"-of", "csv=s=x:p=0",
		Url,
	)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	resolution := strings.TrimSuffix(string(output), "\n")

	return resolution, nil
}

func (r *FilesService) ProcessVideo(InputUrl string, OutputPath string, Resolution string) error {
	if err := os.MkdirAll(OutputPath, os.FileMode(0522)); err != nil {
		return err
	}

	segmentFile := fmt.Sprintf("%s/segment%%d.ts", OutputPath)
	indexFile := fmt.Sprintf("%s/index.m3u8", OutputPath)

	cmd := exec.Command(
		"ffmpeg",
		"-i", InputUrl,
		"-vf", fmt.Sprintf("scale=%s", Resolution),
		"-preset", "slow",
		"-crf", "18",
		"-hls_time", "10",
		"-hls_list_size", "0",
		"-hls_segment_filename",
		segmentFile,
		"-f", "hls",
		indexFile,
	)

	return cmd.Run()
}
