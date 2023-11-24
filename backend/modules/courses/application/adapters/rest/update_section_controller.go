package rest

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/adapters/rest/mappers"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/usecases"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/shared/application/middlewares"
)

type UpdateSectionData struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	AvaliationUrl string `json:"avaliation_url"`
}

func (s *CreateCourseController) UpdateSection(ctx *gin.Context) {
	var updateSectionForm UpdateSectionData

	user := ctx.MustGet("user").(middlewares.User)
	if err := ctx.ShouldBindJSON(&updateSectionForm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	course, err := s.updateSectionUseCase.Execute(
		usecases.UpdateSectionDTO{
			Name:          updateSectionForm.Name,
			Description:   updateSectionForm.Description,
			UserId:        user.Id,
			SectionId:     ctx.Param("id"),
			AvaliationUrl: updateSectionForm.AvaliationUrl,
		},
	)

	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(200, mappers.CourseToGinH(course))

}
