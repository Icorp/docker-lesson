package main

import (
	"github.com/Icorp/docker-lesson/http_server/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/calc")
	{
		v1.GET("/add/:a/:b", controllers.AddPoints)
		v1.GET("/sub/:a/:b", controllers.SubtractPoints)
		v1.GET("/mul/:a/:b", controllers.MultiplyPoints)
		v1.GET("/div/:a/:b", controllers.DividePoints)
	}

	err := router.Run()
	if err != nil {
		panic(err)
	}
}
