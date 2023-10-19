package ports

import "github.com/matheusvmallmann/plataforma-ead/service-course/domain/entities"

type GetFilters struct {
	Status string
}

type VideosRepository interface {
	Create(video *entities.Video) error
	Find(id string) (*entities.Video, error)
	Get(filters GetFilters) ([]*entities.Video, error)
	Update(video *entities.Video) error
	Delete(id string) error
}
