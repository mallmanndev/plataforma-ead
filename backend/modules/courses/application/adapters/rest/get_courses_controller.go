package rest

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/application/adapters/rest/mappers"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/ports"
)

func (c *CreateCourseController) GetCourses(ctx *gin.Context) {
	id := ctx.Query("id")
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
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	coursesResponse := make([]gin.H, len(courses))
	for i, item := range courses {
		coursesResponse[i] = mappers.CourseToGinH(item)
	}

	ctx.JSON(http.StatusOK, coursesResponse)
}
