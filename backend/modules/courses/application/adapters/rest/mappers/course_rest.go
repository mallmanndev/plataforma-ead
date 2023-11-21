package mappers

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/entities"
)

func CourseToGinH(c *entities.Course) gin.H {
	sectionsData := make([]gin.H, len(c.Sections()))
	for i, section := range c.Sections() {
		sectionsData[i] = CourseSectionToGinH(section)
	}
	return gin.H{
		"id":          c.Id(),
		"name":        c.Name(),
		"description": c.Description(),
		"createdAt":   c.CreatedAt(),
		"updatedAt":   c.UpdatedAt(),
		"sections":    sectionsData,
	}
}

func CourseSectionToGinH(cs *entities.CourseSection) gin.H {
	itemsData := make([]gin.H, len(cs.Itens()))
	for i, item := range cs.Itens() {
		itemsData[i] = CourseItemToGinH(item)
	}
	return gin.H{
		"id":          cs.Id(),
		"name":        cs.Name(),
		"description": cs.Description(),
		"createdAt":   cs.CreatedAt(),
		"updatedAt":   cs.UpdatedAt(),
		"itens":       itemsData,
	}
}

func CourseItemToGinH(ci *entities.CourseItem) gin.H {
	return gin.H{
		"id":          ci.Id(),
		"title":       ci.Title(),
		"description": ci.Description(),
		"createdAt":   ci.CreatedAt(),
		"updatedAt":   ci.UpdatedAt(),
	}
}
