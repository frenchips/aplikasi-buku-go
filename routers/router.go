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
		api.GET("/categories", controller.GetAllCategory)
		api.GET("/categories/:id/books", controller.GetBooksByCategory)

		// BUKU
		api.POST("/books", controller.InsertBuku)
		api.GET("/books/:id", controller.GetBuku)
		api.DELETE("/books/:id", controller.DeleteBuku)
		api.GET("/books", controller.GetAllBooks)
	}

	return r
}
