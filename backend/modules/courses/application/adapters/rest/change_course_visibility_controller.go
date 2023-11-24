package rest

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/shared/application/middlewares"
)

type ChangeVisibilityData struct {
	Visibility string `json:"visibility"`
}

func (c *CreateCourseController) ChangeCourseVisibility(ctx *gin.Context) {
	var ChangeVisibilityData ChangeVisibilityData

	user := ctx.MustGet("user").(middlewares.User)
	if err := ctx.ShouldBindJSON(&ChangeVisibilityData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if ChangeVisibilityData.Visibility == "public" {
		if err := c.makeCourseVisible.Execute(ctx.Param("id"), user.Id); err != nil {
			log.Println(err)
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
	} else {
		if err := c.makeCourseInvisible.Execute(ctx.Param("id"), user.Id); err != nil {
			log.Println(err)
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
	}

	ctx.JSON(200, gin.H{"message": "Course updated successfully."})

}
