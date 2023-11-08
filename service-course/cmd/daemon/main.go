package main

import (
	"log"
	"time"

	"github.com/matheusvmallmann/plataforma-ead/service-course/application/adapters/repositories"
	"github.com/matheusvmallmann/plataforma-ead/service-course/application/adapters/services"
	"github.com/matheusvmallmann/plataforma-ead/service-course/application/usecases"
	"github.com/matheusvmallmann/plataforma-ead/service-course/utils"
)

func main() {
	db, disconnect := utils.GetDb("dev")
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
