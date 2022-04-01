package main

import (
	"github.com/Icorp/docker-lesson/http_server/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/calc")
	{
		v1.GET("/add/:a/:b", routes.AddPoints)
		v1.GET("/sub/:a/:b", routes.SubtractPoints)
		v1.GET("/mul/:a/:b", routes.MultiplyPoints)
		v1.GET("/div/:a/:b", routes.DividePoints)
	}

	err := router.Run()
	if err != nil {
		panic(err)
	}
}
