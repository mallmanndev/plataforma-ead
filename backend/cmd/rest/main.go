package main

import (
	"github.com/gin-gonic/gin"
	courses "github.com/matheusvmallmann/plataforma-ead/backend/modules/courses"
	users "github.com/matheusvmallmann/plataforma-ead/backend/modules/users"
	"github.com/matheusvmallmann/plataforma-ead/backend/utils"
)

func main() {
	db, disconnect := utils.GetDb()
	defer disconnect()

	r := gin.Default()

	users.Routes(&r.RouterGroup, db)
	courses.Routes(&r.RouterGroup, db)

	r.Run(":3000")
}
