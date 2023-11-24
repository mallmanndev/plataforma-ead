package main

import (
	"log"
	"time"

	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/adapters/repositories"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/adapters/services"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/usecases"
	"github.com/matheusvmallmann/plataforma-ead/backend/utils"
)

func main() {
	db, disconnect := utils.GetDb()
	defer disconnect()
	videosRepository := repositories.NewVideosRepository(db)
	filesService := services.NewFilesService()
	useCase := usecases.NewProcessVideoTwo(videosRepository, filesService)

	intervalTime := 5

	log.Printf("Video daemon started. Runing in: %d", intervalTime)

	for {
		useCase.Execute()
		time.Sleep(time.Duration(intervalTime) * time.Second)
	}
}
