package rest

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/adapters/rest/mappers"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/usecases"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/shared/application/middlewares"
)

type UpdateItemData struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (s *CreateCourseController) UpdateItem(ctx *gin.Context) {
	var updateItemData UpdateItemData

	user := ctx.MustGet("user").(middlewares.User)
	if err := ctx.ShouldBindJSON(&updateItemData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course, err := s.updateItemUseCase.Execute(
		usecases.UpdateItemInput{
			Id:          updateItemData.Id,
			Title:       updateItemData.Title,
			Description: updateItemData.Description,
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
