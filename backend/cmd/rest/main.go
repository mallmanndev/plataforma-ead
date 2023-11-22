package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/matheusvmallmann/plataforma-ead/backend/docs"
	courses "github.com/matheusvmallmann/plataforma-ead/backend/modules/courses"
	users "github.com/matheusvmallmann/plataforma-ead/backend/modules/users"
	"github.com/matheusvmallmann/plataforma-ead/backend/utils"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

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
