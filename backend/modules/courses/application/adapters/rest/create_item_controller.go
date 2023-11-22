package rest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/adapters/rest/mappers"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/usecases"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/shared/application/middlewares"
)

type CreateItemData struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	SectionId   string `json:"section_id"`
	VideoId     string `json:"video_id"`
}

func (c *CreateCourseController) CreateItem(ctx *gin.Context) {
	var createItemData CreateItemData

	user := ctx.MustGet("user").(middlewares.User)
	if err := ctx.ShouldBindJSON(&createItemData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(createItemData)
	course, err := c.createItemUseCase.Execute(
		usecases.CreateItemInput{
			Title:       createItemData.Title,
			Description: createItemData.Description,
			SectionId:   createItemData.SectionId,
			VideoId:     createItemData.VideoId,
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
