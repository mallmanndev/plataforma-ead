package ports

import "github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"

type VideosRepository interface {
	Create(video *entities.Video) error
	// Find(id string) (*entities.Video, error)
}
