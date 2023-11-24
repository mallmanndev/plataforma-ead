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
	DiscordUrl  string `json:"discord_url"`
}

func (c *CreateCourseController) UpdateCourse(ctx *gin.Context) {
	var updateCourseData UpdateCourseData

	user := ctx.MustGet("user").(middlewares.User)
	if err := ctx.ShouldBindJSON(&updateCourseData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	course, err := c.updateCourseUseCase.Execute(
		usecases.UpdateCourseUseCaseDTO{
			Id:          ctx.Param("id"),
			Name:        updateCourseData.Name,
			Description: updateCourseData.Description,
			UserId:      user.Id,
			DiscordUrl:  updateCourseData.DiscordUrl,
		},
	)

	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, mappers.CourseToGinH(course))

}
