package rest

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/adapters/rest/mappers"
)

func (c *CreateCourseController) GetSection(ctx *gin.Context) {
	course, err := c.coursesRepo.FindBySectionId(ctx.Param("id"))

	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	if course == nil {
		ctx.JSON(404, gin.H{"message": "Course not found"})
		return
	}

	section := course.FindSection(ctx.Param("id"))

	ctx.JSON(200, mappers.CourseSectionToGinH(section))
}
