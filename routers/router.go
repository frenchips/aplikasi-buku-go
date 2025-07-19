package routers

import (
	"aplikasi-buku-go/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		// TO DO ROUTER
		// CATEGORY
		// api.GET("/helloWorld", HelloWorld)
		api.POST("/categories", controller.InsertCategory)
		api.PUT("/categories/:id", controller.UpdateCategory)
		api.DELETE("/categories/:id", controller.DeleteCategory)
	}

	return r
}
