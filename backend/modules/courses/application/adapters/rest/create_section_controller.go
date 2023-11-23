package rest

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/adapters/rest/mappers"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/usecases"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/shared/application/middlewares"
)

type CreateSectionData struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	CourseId      string `json:"course_id"`
	AvaliationUrl string `json:"avaliation_url"`
}

func (c *CreateCourseController) CreateSection(ctx *gin.Context) {
	var createSectionForm CreateSectionData

	user := ctx.MustGet("user").(middlewares.User)
	if err := ctx.ShouldBindJSON(&createSectionForm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course, err := c.createSectionUseCase.Execute(
		usecases.CreateSectionDTO{
			Name:          createSectionForm.Name,
			Description:   createSectionForm.Description,
			CourseId:      createSectionForm.CourseId,
			UserId:        user.Id,
			AvaliationUrl: createSectionForm.AvaliationUrl,
		},
	)

	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, mappers.CourseToGinH(course))

}
