package users

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/users/aplication/adapters/rest"
	"go.mongodb.org/mongo-driver/mongo"
)

func Routes(r *gin.RouterGroup, db *mongo.Database) {
	controller := rest.NewUserController(db)

	r.POST("/login", controller.Login)
	r.POST("/user", controller.CreateAccount)
}
