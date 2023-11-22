package rest

import (
	"io"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/matheusvmallmann/plataforma-ead/backend/modules/shared/application/middlewares"
)

func (c *CreateCourseController) VideoUpload(ctx *gin.Context) {
	user := ctx.MustGet("user").(middlewares.User)
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(400, gin.H{"message": "Error to get file"})
		return
	}

	fileExtension := strings.Replace(filepath.Ext(file.Filename), ".", "", 1)
	upload, err := c.videoUploadUseCase.CreateFile(
		fileExtension,
		file.Size,
		user.Id,
	)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	openedFile, err := file.Open()
	if err != nil {
		ctx.JSON(400, gin.H{"message": "Error to open file"})
		return
	}
	bytes, err := io.ReadAll(openedFile)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "Error to read file"})
		return
	}

	err = upload.SendChunk(bytes)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "Error to send chunk"})
		return
	}

	video, err := upload.Execute()
	if err != nil {
		ctx.JSON(400, gin.H{"message": "Error to execute upload"})
		return
	}

	ctx.JSON(200, gin.H{
		"id": video.Id(),
	})
}
