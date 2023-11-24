package mappers

import (
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/adapters/repositories/models"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/apptimer"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/entities"
)

func VideoModelToVideoEntity(Model models.VideoModel) *entities.Video {
	videoEntity := entities.NewCompleteVideo(
		apptimer.NewAppTimer(),
		Model.Id,
		Model.Type,
		Model.TmpUrl,
		Model.Status,
		Model.Duration,
		Model.Size,
		Model.UserId,
		Model.CreatedAt,
		Model.UpdatedAt,
		Model.Url,
	)

	for _, res := range Model.Resolutions {
		videoEntity.AddResolution(entities.VideoResolution{
			Resolution: res.Resolution,
			URL:        res.URL,
		})
	}

	return videoEntity
}
