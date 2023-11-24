package mappers

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/courses/domain/entities"
)

func CourseToGinH(c *entities.Course) gin.H {
	var discord_url string
	if c.DiscordUrl() != nil {
		discord_url = c.DiscordUrl().String()
	}

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
		"discord_url": discord_url,
	}
}

func CourseSectionToGinH(cs *entities.CourseSection) gin.H {
	var avaliation_url string
	if cs.Avaliation() != nil {
		avaliation_url = cs.Avaliation().String()
	}

	itemsData := make([]gin.H, len(cs.Itens()))
	for i, item := range cs.Itens() {
		itemsData[i] = CourseItemToGinH(item)
	}
	return gin.H{
		"id":             cs.Id(),
		"name":           cs.Name(),
		"description":    cs.Description(),
		"createdAt":      cs.CreatedAt(),
		"updatedAt":      cs.UpdatedAt(),
		"itens":          itemsData,
		"avaliation_url": avaliation_url,
	}
}

func CourseItemToGinH(ci *entities.CourseItem) gin.H {
	return gin.H{
		"id":          ci.Id(),
		"title":       ci.Title(),
		"description": ci.Description(),
		"videoId":     ci.VideoId(),
		"createdAt":   ci.CreatedAt(),
		"updatedAt":   ci.UpdatedAt(),
	}
}
