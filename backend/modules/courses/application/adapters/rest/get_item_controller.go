package rest

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/adapters/rest/mappers"
)

func (c *CreateCourseController) GetItem(ctx *gin.Context) {
	course, err := c.coursesRepo.FindByItemId(ctx.Param("id"))

	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if course == nil {
		ctx.JSON(404, gin.H{"error": "Course not found"})
		return
	}

	item, _ := course.FindItem(ctx.Param("id"))

	ctx.JSON(200, mappers.CourseItemToGinH(item))
}
