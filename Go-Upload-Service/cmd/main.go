package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kalyanKumarPokkula/vercel/controllers"
)

func main() {
	r := gin.Default()

	r.POST("/deploy", controllers.Deploy)

	r.Run() // listen and serve on 0.0.0.0:8080
}