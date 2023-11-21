package rest

import (
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/adapters/repositories"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/usecases"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/ports"
	"go.mongodb.org/mongo-driver/mongo"
)

type CreateCourseController struct {
	coursesRepo         ports.CourseRepository
	createCourseUseCase *usecases.CreateCourseUseCase
}

func NewCourseServer(db *mongo.Database) *CreateCourseController {
	coursesRepo := repositories.NewCourseRepositories(db)
	createCourseUseCase := usecases.NewCreateCourseUseCase(coursesRepo)

	return &CreateCourseController{
		coursesRepo:         coursesRepo,
		createCourseUseCase: createCourseUseCase,
	}
}
