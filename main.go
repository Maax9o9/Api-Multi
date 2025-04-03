package main

import (
    infresHouse "Multi/src/house/infrestructure"
    routesHouse "Multi/src/house/infrestructure/routes"
    infresNotification "Multi/src/notification/infrestructure"
    routesNotification "Multi/src/notification/infrestructure/routes"
    infresUser "Multi/src/user/infrestructure"
    routesUser "Multi/src/user/infrestructure/routes"
    infresWeather "Multi/src/weather"
    routesWeather "Multi/src/weather/infrestructure/routes"
    controllerWeather "Multi/src/weather/infrestructure/controllers"
    infresLight "Multi/src/interruptors/light/infrestructure"
    routesLight "Multi/src/interruptors/light/infrestructure/routes"
    controllerLight "Multi/src/interruptors/light/infrestructure/controllers"
    infresLightOutside "Multi/src/interruptors/lightOutside/infrestructure"
    routesLightOutside "Multi/src/interruptors/lightOutside/infrestructure/routes"
    controllerLightOutside "Multi/src/interruptors/lightOutside/infrestructure/controllers"
    infresMovement "Multi/src/interruptors/movement/infrestructure"
    routesMovement "Multi/src/interruptors/movement/infrestructure/routes"
    controllerMovement "Multi/src/interruptors/movement/infrestructure/controllers"
    infresDoor "Multi/src/interruptors/door/infrestructure"
    routesDoor "Multi/src/interruptors/door/infrestructure/routes"
    controllersDoor "Multi/src/interruptors/door/infrestructure/controllers"
    infresWindow "Multi/src/interruptors/window/infrestructure"
    routesWindow "Multi/src/interruptors/window/infrestructure/routes"
    controllerWindow "Multi/src/interruptors/window/infrestructure/controllers"
    infresGas "Multi/src/interruptors/gas/infrestructure"
    routesGas "Multi/src/interruptors/gas/infrestructure/routes"
    controllerGas "Multi/src/interruptors/gas/infrestructure/controllers"
    "log"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.Use(cors.Default())

    alertWeatherService, receiveWeatherService, rabbitMQWeather := infresWeather.InitWeather()
    alertWeatherController := controllerWeather.NewAlertWeatherController(alertWeatherService)
    receiveWeatherAllController := controllerWeather.NewReceiveWeatherController(receiveWeatherService)
    receiveWeatherByIDController := controllerWeather.NewReceiveWeatherByIDController(receiveWeatherService)
    routesWeather.WeatherRoutes(router, alertWeatherController, receiveWeatherAllController, receiveWeatherByIDController)

    createUserController, showAllUsersController, showUserByIDController, loginController := infresUser.InitUser()
    routesUser.UserRoutes(router, createUserController, showAllUsersController, showUserByIDController, loginController)

    addHouseController, showAllHousesController, showHouseByIDController, editHouseController := infresHouse.InitHouse()
    routesHouse.HouseRoutes(router, addHouseController, showAllHousesController, showHouseByIDController, editHouseController)

    createNotificationController, showAllNotificationsController, showNotificationByIDController := infresNotification.InitNotification()
    routesNotification.NotificationRoutes(router, createNotificationController, showAllNotificationsController, showNotificationByIDController)

    alertLightService, receiveLightService, rabbitMQLight := infresLight.InitLight()
    alertLightController := controllerLight.NewAlertLightController(alertLightService)
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
    alertDoorController := controllersDoor.NewAlertDoorController(alertDoorService)
    reciveDoorController := controllersDoor.NewGetAllDoorController(receiveDoorService)
    reciveDoorByIDController := controllersDoor.NewGetDoorByIDController(receiveDoorService)
    routesDoor.DoorRoutes(router, alertDoorController, reciveDoorController, reciveDoorByIDController)

    alertWindowService, receiveWindowService, rabbitMQWindow := infresWindow.InitWindow()
    alertWindowController := controllerWindow.NewAlertWindowController(alertWindowService)
    receiveWindowAllController := controllerWindow.NewReceiveWindowController(receiveWindowService)
    receiveWindowByIDController := controllerWindow.NewReceiveWindowByIDController(receiveWindowService)
    routesWindow.WindowRoutes(router, alertWindowController, receiveWindowAllController, receiveWindowByIDController)

    alertGasService, receiveGasService, rabbitMQGas := infresGas.InitGas()
    alertGasController := controllerGas.NewAlertGasController(alertGasService)
    receiveGasAllController := controllerGas.NewReceiveGasController(receiveGasService)
    receiveGasByIDController := controllerGas.NewReceiveGasByIDController(receiveGasService)
    routesGas.GasRoutes(router, alertGasController, receiveGasAllController, receiveGasByIDController)

    defer rabbitMQGas.Close()
    defer rabbitMQWeather.Close()
    defer rabbitMQMovement.Close()
    defer rabbitMQLight.Close()
    defer rabbitMQLightOutside.Close()
    defer rabbitMQDoor.Close()
    defer rabbitMQWindow.Close()

    // Iniciar servidor
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}