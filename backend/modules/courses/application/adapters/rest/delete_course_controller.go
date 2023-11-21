package rest

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/usecases"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/shared/application/middlewares"
)

func (c *CreateCourseController) DeleteCourse(ctx *gin.Context) {
	user := ctx.MustGet("user").(middlewares.User)

	err := c.deleteCourseUseCase.Execute(
		usecases.DeleteCourseUseCaseDataDTO{
			Id:     ctx.Param("id"),
			UserId: user.Id,
		},
	)

	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Course deleted successfully."})

}
