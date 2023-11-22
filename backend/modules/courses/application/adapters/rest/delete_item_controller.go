package rest

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/usecases"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/shared/application/middlewares"
)

func (s *CreateCourseController) DeleteItem(ctx *gin.Context) {
	user := ctx.MustGet("user").(middlewares.User)

	_, err := s.deleteItemUseCase.Execute(
		usecases.DeleteItemInput{
			UserId: user.Id,
			Id:     ctx.Param("id"),
		},
	)

	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Section deleted successfully"})

}
