package courses

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/adapters/rest"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/shared/application/middlewares"
	"go.mongodb.org/mongo-driver/mongo"
)

func Routes(r *gin.RouterGroup, db *mongo.Database) {
	controller := rest.NewCourseServer(db)

	r.POST("/course", middlewares.VerifyTokenMiddleware, controller.CreateCourse)
	r.GET("/courses", middlewares.VerifyTokenMiddleware, controller.GetCourses)
	r.GET("/courses/:id", middlewares.VerifyTokenMiddleware, controller.GetCourse)
	r.PUT("/courses/:id", middlewares.VerifyTokenMiddleware, controller.UpdateCourse)
	r.DELETE("/courses/:id", middlewares.VerifyTokenMiddleware, controller.DeleteCourse)
	r.PATCH("/courses/:id/change-visibility", middlewares.VerifyTokenMiddleware, controller.ChangeCourseVisibility)
	r.POST("/sections", middlewares.VerifyTokenMiddleware, controller.CreateSection)
	r.PUT("/sections/:id", middlewares.VerifyTokenMiddleware, controller.UpdateSection)
	r.DELETE("/sections/:id", middlewares.VerifyTokenMiddleware, controller.DeleteSection)
	r.GET("/sections/:id", middlewares.VerifyTokenMiddleware, controller.GetSection)
}
