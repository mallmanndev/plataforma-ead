package services

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/ports"
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
	os.MkdirAll(filepath.Dir(File.Url), 0770)

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

func (r *FilesService) Delete(Url string) error {
	return os.Remove(Url)
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

	var outbuf, errbuf bytes.Buffer
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	log.Printf("Running with user id: %d, group id: %d", os.Getuid(), os.Getgid())

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Credential: &syscall.Credential{
			Uid: uint32(os.Getuid()),
			Gid: uint32(os.Getgid()),
		},
	}

	if err := cmd.Run(); err != nil {
		log.Println("Terminal error", errbuf.String())
		return err
	}
	return nil
}
