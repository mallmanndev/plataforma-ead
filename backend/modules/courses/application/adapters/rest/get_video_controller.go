package rest

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (c *CreateCourseController) GetVideo(ctx *gin.Context) {
	video, err := c.videosRepo.Find(ctx.Param("id"))

	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if video == nil {
		ctx.JSON(404, gin.H{"error": "Video not found"})
		return
	}

	var resolutions []gin.H
	for _, res := range video.GetResolutions() {
		resolutions = append(resolutions, gin.H{
			"resolution": res.Resolution,
			"url":        res.URL,
		})
	}

	ctx.JSON(200, gin.H{
		"id":          video.Id(),
		"type":        video.Type(),
		"status":      video.Status(),
		"size":        video.Size(),
		"createdAt":   video.CreatedAt().String(),
		"updatedAt":   video.UpdatedAt().String(),
		"url":         video.Url(),
		"resolutions": resolutions,
	})

}
