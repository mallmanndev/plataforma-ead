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
}
