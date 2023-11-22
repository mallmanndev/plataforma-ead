package courses

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/adapters/rest"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/shared/application/middlewares"
	"go.mongodb.org/mongo-driver/mongo"
)

func Routes(r *gin.RouterGroup, db *mongo.Database) {
	controller := rest.NewCourseServer(db)

	r.Static("/files/videos", "/videos")

	// COURSES
	r.POST("/courses", middlewares.VerifyTokenMiddleware, controller.CreateCourse)
	r.GET("/courses", middlewares.VerifyTokenMiddleware, controller.GetCourses)
	r.GET("/courses/:id", middlewares.VerifyTokenMiddleware, controller.GetCourse)
	r.PUT("/courses/:id", middlewares.VerifyTokenMiddleware, controller.UpdateCourse)
	r.DELETE("/courses/:id", middlewares.VerifyTokenMiddleware, controller.DeleteCourse)
	r.PATCH("/courses/:id/change-visibility", middlewares.VerifyTokenMiddleware, controller.ChangeCourseVisibility)

	// SECTIONS
	r.POST("/sections", middlewares.VerifyTokenMiddleware, controller.CreateSection)
	r.PUT("/sections/:id", middlewares.VerifyTokenMiddleware, controller.UpdateSection)
	r.DELETE("/sections/:id", middlewares.VerifyTokenMiddleware, controller.DeleteSection)
	r.GET("/sections/:id", middlewares.VerifyTokenMiddleware, controller.GetSection)

	// ITENS
	r.POST("/itens", middlewares.VerifyTokenMiddleware, controller.CreateItem)
	r.PUT("/itens/:id", middlewares.VerifyTokenMiddleware, controller.UpdateItem)
	r.DELETE("/itens/:id", middlewares.VerifyTokenMiddleware, controller.DeleteItem)
	r.GET("/itens/:id", middlewares.VerifyTokenMiddleware, controller.GetItem)

	// VIDEOS
	r.GET("/videos/:id", middlewares.VerifyTokenMiddleware, controller.GetVideo)
	r.POST("/videos", middlewares.VerifyTokenMiddleware, controller.VideoUpload)
}
