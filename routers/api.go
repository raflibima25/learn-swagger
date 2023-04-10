package routers

import (
	"DTS-Materi/learn-swagger-materi/controllers"

	"github.com/gin-gonic/gin"

	_ "DTS-Materi/learn-swagger-materi/docs"

	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerFiles "github.com/swaggo/files"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	// Read
	router.GET("/cars/:id", controllers.GetOneCars)
	// Create
	router.POST("/cars", controllers.CreateCars)
	// Read All
	router.GET("/cars", controllers.GetAllCars)
	// Update
	router.PATCH("/cars/:id", controllers.UpdateCars)
	// Delete
	router.DELETE("/cars/:id", controllers.DeleteCars)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
