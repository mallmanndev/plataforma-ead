package rest

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/adapters/rest/mappers"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/usecases"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/shared/application/middlewares"
)

type UpdateCourseData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c *CreateCourseController) UpdateCourse(ctx *gin.Context) {
	var createCourseForm CreateCourseData

	user := ctx.MustGet("user").(middlewares.User)
	if err := ctx.ShouldBindJSON(&createCourseForm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course, err := c.updateCourseUseCase.Execute(
		usecases.UpdateCourseUseCaseDTO{
			Id:          ctx.Param("id"),
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

	ctx.JSON(200, mappers.CourseToGinH(course))

}
