package services

import (
	"github.com/matheusvmallmann/plataforma-ead/service-course/domain/ports"
	"os"
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

func (r *FilesService) Close() error {
	return r.file.Close()
}

func (r *FilesService) Remove() error {
	return os.Remove(r.url)
}
