package main

import (
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
	useCase := usecases.NewProcessVideo(videosRepository, filesService)

	useCase.Execute()
}
