package rest

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/adapters/rest/mappers"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/ports"
)

func (c *CreateCourseController) GetCourse(ctx *gin.Context) {
	id := ctx.Param("id")
	userID := ctx.Query("userId")
	visibleStr := ctx.Query("visible")
	visible, _ := strconv.ParseBool(visibleStr)

	courses, err := c.coursesRepo.Get(ports.GetCourseFilters{
		Id:      id,
		UserId:  userID,
		Visible: visible,
	})

	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	if len(courses) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Course not found"})
		return
	}

	ctx.JSON(http.StatusOK, mappers.CourseToGinH(courses[0]))
}
