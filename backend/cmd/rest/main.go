package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/matheusvmallmann/plataforma-ead/backend/docs"
	courses "github.com/matheusvmallmann/plataforma-ead/backend/modules/courses"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/adapters/repositories"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/adapters/services"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/usecases"
	users "github.com/matheusvmallmann/plataforma-ead/backend/modules/users"
	"github.com/matheusvmallmann/plataforma-ead/backend/utils"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
)

func VideoProcessDaemon(db *mongo.Database) {
	videosRepo := repositories.NewVideosRepository(db)
	filesService := services.NewFilesService()

	useCase := usecases.NewProcessVideoTwo(videosRepo, filesService)

	intervalTime := 5
	log.Printf("Video daemon started. Runing in: %d", intervalTime)

	for {
		useCase.Execute()
		time.Sleep(time.Duration(intervalTime) * time.Second)
	}
}

// @title Plataforma EAD
// @version 1.0
// @description Descrição da sua API
// @termsOfService https://example.com/terms/
// @contact.name API Support
// @contact.url https://www.example.com/support
// @host localhost:8080
// @BasePath /v1
func main() {
	db, disconnect := utils.GetDb()
	defer disconnect()

	if os.Getenv("RUN_DAEMON") == "true" {
		go VideoProcessDaemon(db)
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	users.Routes(&r.RouterGroup, db)
	courses.Routes(&r.RouterGroup, db)

	r.Run(":3000")
}
