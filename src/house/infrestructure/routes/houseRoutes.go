package routes

import (
	"Multi/src/house/infrestructure/controllers"
	"Multi/src/user/infrestructure/middlewares"

	"github.com/gin-gonic/gin"
)

// SetupHouseRoutes configura todas las rutas relacionadas con casas
func SetupHouseRoutes(
	router *gin.Engine,
	createHouseController *controllers.CreateHouseController,
	getHouseByIDController *controllers.GetHouseByIDController,
	getUserHousesController *controllers.GetUserHousesController,
	updateHouseController *controllers.UpdateHouseController,
	deleteHouseController *controllers.DeleteHouseController,
	updateLocationController *controllers.UpdateLocationController,
	uploadImageController *controllers.UploadImageController,
	addMemberToHouseController *controllers.AddMemberToHouseController,
) {
	houseRoutes := router.Group("/api/houses")
	houseRoutes.Use(middlewares.AuthMiddleware())
	{
		// Rutas básicas CRUD
		houseRoutes.POST("", createHouseController.Handle)
		houseRoutes.GET("", getUserHousesController.Handle)
		houseRoutes.GET("/:id", getHouseByIDController.Handle)
		houseRoutes.PUT("/:id", updateHouseController.Handle)
		houseRoutes.DELETE("/:id", deleteHouseController.Handle)

		// Rutas específicas
		houseRoutes.PUT("/:id/location", updateLocationController.Handle)
		houseRoutes.POST("/images", uploadImageController.Handle)

		// Rutas de miembros
		houseRoutes.POST("/:id/members", addMemberToHouseController.Handle)
	}

	// Configuración de archivos estáticos para acceder a las imágenes subidas
	router.Static("/uploads", "./uploads")
}
