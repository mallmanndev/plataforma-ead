package rest

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/usecases"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/shared/application/middlewares"
)

type CreateCourseData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c *CreateCourseController) CreateCourse(ctx *gin.Context) {
	var createCourseForm CreateCourseData

	user := ctx.MustGet("user").(middlewares.User)
	if err := ctx.ShouldBindJSON(&createCourseForm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course, err := c.createCourseUseCase.Execute(
		usecases.CreateCourseUseCaseDTO{
			Name:        createCourseForm.Name,
			Description: createCourseForm.Description,
			UserId:      user.Id,
		},
	)

	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	log.Println(course)
	ctx.JSON(200, gin.H{
		"id":          course.Id(),
		"name":        course.Name(),
		"description": course.Description(),
		"visible":     course.IsVisible(),
		"sections":    course.Sections(),
		"created_at":  course.CreatedAt(),
		"updated_at":  course.UpdatedAt(),
	})

}
