package main

import (
	infresHallDoor "Multi/src/hallDoor/infrestructure"
	controllerHallDoor "Multi/src/hallDoor/infrestructure/controllers"
	routesHallDoor "Multi/src/hallDoor/infrestructure/routes"
	infresHallWindow "Multi/src/hallWindow/infrestructure"
	controllerHallWindow "Multi/src/hallWindow/infrestructure/controllers"
	routesHallWindow "Multi/src/hallWindow/infrestructure/routes"
	infresHouse "Multi/src/house/infrestructure"
	routesHouse "Multi/src/house/infrestructure/routes"
	infresDoor "Multi/src/interruptors/door/infrestructure"
	controllersDoor "Multi/src/interruptors/door/infrestructure/controllers"
	routesDoor "Multi/src/interruptors/door/infrestructure/routes"
	infresGas "Multi/src/interruptors/gas/infrestructure"
	controllerGas "Multi/src/interruptors/gas/infrestructure/controllers"
	routesGas "Multi/src/interruptors/gas/infrestructure/routes"
	infresLight "Multi/src/interruptors/light/infrestructure"
	controllerLight "Multi/src/interruptors/light/infrestructure/controllers"
	routesLight "Multi/src/interruptors/light/infrestructure/routes"
	infresLightOutside "Multi/src/interruptors/lightOutside/infrestructure"
	controllerLightOutside "Multi/src/interruptors/lightOutside/infrestructure/controllers"
	routesLightOutside "Multi/src/interruptors/lightOutside/infrestructure/routes"
	infresMovement "Multi/src/interruptors/movement/infrestructure"
	controllerMovement "Multi/src/interruptors/movement/infrestructure/controllers"
	routesMovement "Multi/src/interruptors/movement/infrestructure/routes"
	infresWindow "Multi/src/interruptors/window/infrestructure"
	controllerWindow "Multi/src/interruptors/window/infrestructure/controllers"
	routesWindow "Multi/src/interruptors/window/infrestructure/routes"
	infresNotification "Multi/src/notification/infrestructure"
	routesNotification "Multi/src/notification/infrestructure/routes"
	infresUser "Multi/src/user/infrestructure"
	routesUser "Multi/src/user/infrestructure/routes"
	infresWeather "Multi/src/weather"
	controllerWeather "Multi/src/weather/infrestructure/controllers"
	routesWeather "Multi/src/weather/infrestructure/routes"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	alertWeatherService, receiveWeatherService, rabbitMQWeather := infresWeather.InitWeather()
	alertWeatherController := controllerWeather.NewAlertWeatherController(alertWeatherService)
	receiveWeatherAllController := controllerWeather.NewReceiveWeatherController(receiveWeatherService)
	receiveWeatherByIDController := controllerWeather.NewReceiveWeatherByIDController(receiveWeatherService)
	routesWeather.WeatherRoutes(router, alertWeatherController, receiveWeatherAllController, receiveWeatherByIDController)

	createUserController, showAllUsersController, showUserByIDController, loginController := infresUser.InitUser()
	routesUser.UserRoutes(router, createUserController, showAllUsersController, showUserByIDController, loginController)

	houseControllers := infresHouse.InitializeDependencies()
	routesHouse.SetupHouseRoutes(router,
		houseControllers.CreateHouseController,
		houseControllers.GetHouseByIDController,
		houseControllers.GetUserHousesController,
		houseControllers.UpdateHouseController,
		houseControllers.DeleteHouseController,
		houseControllers.UpdateLocationController,
		houseControllers.UploadImageController,
		houseControllers.AddMemberToHouseController,
	)

	createNotificationController, showAllNotificationsController, showNotificationByIDController := infresNotification.InitNotification()
	routesNotification.NotificationRoutes(router, createNotificationController, showAllNotificationsController, showNotificationByIDController)

	alertLightService, receiveLightService, rabbitMQLight := infresLight.InitLight()
	alertLightController := controllerLight.NewAlertLightController(alertLightService, rabbitMQLight)
	receiveLightAllController := controllerLight.NewReceiveLightController(receiveLightService)
	receiveLightByIDController := controllerLight.NewReceiveLightByIDController(receiveLightService)
	routesLight.LightRoutes(router, alertLightController, receiveLightAllController, receiveLightByIDController)

	alertLightOutsideService, receiveLightOutsideService, rabbitMQLightOutside := infresLightOutside.InitLightOutside()
	alertLightOutsideController := controllerLightOutside.NewAlertLightController(alertLightOutsideService)
	receiveLightOutsideAllController := controllerLightOutside.NewReceiveLightController(receiveLightOutsideService)
	receiveLightOutsideByIDController := controllerLightOutside.NewReceiveLightByIDController(receiveLightOutsideService)
	routesLightOutside.LightRoutes(router, alertLightOutsideController, receiveLightOutsideAllController, receiveLightOutsideByIDController)

	alertMovementService, receiveMovementService, rabbitMQMovement := infresMovement.InitMovement()
	alertMovementController := controllerMovement.NewAlertMovementController(alertMovementService)
	receiveMovementAllController := controllerMovement.NewReceiveMovementController(receiveMovementService)
	receiveMovementByIDController := controllerMovement.NewReceiveMovementByIDController(receiveMovementService)
	routesMovement.MovementRoutes(router, alertMovementController, receiveMovementAllController, receiveMovementByIDController)

	alertDoorService, receiveDoorService, rabbitMQDoor := infresDoor.InitDoor()
	alertDoorController := controllersDoor.NewAlertDoorController(alertDoorService, rabbitMQDoor)
	reciveDoorController := controllersDoor.NewGetAllDoorController(receiveDoorService)
	reciveDoorByIDController := controllersDoor.NewGetDoorByIDController(receiveDoorService)
	routesDoor.DoorRoutes(router, alertDoorController, reciveDoorController, reciveDoorByIDController)

	alertWindowService, receiveWindowService, rabbitMQWindow := infresWindow.InitWindow()
	alertWindowController := controllerWindow.NewAlertWindowController(alertWindowService, rabbitMQWindow)
	receiveWindowAllController := controllerWindow.NewReceiveWindowController(receiveWindowService)
	receiveWindowByIDController := controllerWindow.NewReceiveWindowByIDController(receiveWindowService)
	routesWindow.WindowRoutes(router, alertWindowController, receiveWindowAllController, receiveWindowByIDController)

	alertGasService, receiveGasService, rabbitMQGas := infresGas.InitGas()
	alertGasController := controllerGas.NewAlertGasController(alertGasService)
	receiveGasAllController := controllerGas.NewReceiveGasController(receiveGasService)
	receiveGasByIDController := controllerGas.NewReceiveGasByIDController(receiveGasService)
	routesGas.GasRoutes(router, alertGasController, receiveGasAllController, receiveGasByIDController)

	alertHallDoorService, receiveHallDoorService, updateHallDoorService, rabbitMQHallDoor := infresHallDoor.InitHallDoor()
	alertHallDoorController := controllerHallDoor.NewAlertHallDoorController(alertHallDoorService)
	receiveHallDoorAllController := controllerHallDoor.NewReceiveHallDoorController(receiveHallDoorService)
	receiveHallDoorByIDController := controllerHallDoor.NewReceiveHallDoorByIDController(receiveHallDoorService)
	updateHallDoorController := controllerHallDoor.NewUpdateHallDoorController(updateHallDoorService)
	routesHallDoor.HallDoorRoutes(router, alertHallDoorController, receiveHallDoorAllController, receiveHallDoorByIDController, updateHallDoorController)

	alertHallWindowService, receiveHallWindowService, updateHallWindowService, rabbitMQHallWindow := infresHallWindow.InitHallWindow()
	alertHallWindowController := controllerHallWindow.NewAlertHallWindowController(alertHallWindowService)
	receiveHallWindowAllController := controllerHallWindow.NewReceiveHallWindowController(receiveHallWindowService)
	receiveHallWindowByIDController := controllerHallWindow.NewReceiveHallWindowByIDController(receiveHallWindowService)
	updateHallWindowController := controllerHallWindow.NewUpdateHallWindowController(updateHallWindowService)
	routesHallWindow.HallWindowRoutes(router, alertHallWindowController, receiveHallWindowAllController, receiveHallWindowByIDController, updateHallWindowController)

	defer rabbitMQGas.Close()
	defer rabbitMQWeather.Close()
	defer rabbitMQMovement.Close()
	defer rabbitMQLight.Close()
	defer rabbitMQLightOutside.Close()
	defer rabbitMQDoor.Close()
	defer rabbitMQWindow.Close()
	defer rabbitMQHallDoor.Close()
	defer rabbitMQHallWindow.Close()

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
