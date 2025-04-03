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
    receiveWeatherController := controllerWeather.NewReceiveWeatherController(receiveWeatherService)
    routesWeather.WeatherRoutes(router, alertWeatherController, receiveWeatherController)

    createUserController, showUserController, loginController := infresUser.InitUser()
    routesUser.UserRoutes(router, createUserController, showUserController, loginController)

    addHouseController, showHouseController, editHouseController := infresHouse.InitHouse()
    routesHouse.HouseRoutes(router, addHouseController, showHouseController, editHouseController)

    createNotificationController, showNotificationController := infresNotification.InitNotification()
    routesNotification.NotificationRoutes(router, createNotificationController, showNotificationController)

    alertLightService, receiveLightService, rabbitMQLight := infresLight.InitLight()
    alertLightController := controllerLight.NewAlertLightController(alertLightService)
    receiveLightController := controllerLight.NewReceiveLightController(receiveLightService)
    routesLight.LightRoutes(router, alertLightController, receiveLightController)

    alertMovementService, receiveMovementService, rabbitMQMovement := infresMovement.InitMovement()
    alertMovementController := controllerMovement.NewAlertMovementController(alertMovementService)
    receiveMovementController := controllerMovement.NewReceiveMovementController(receiveMovementService)
    routesMovement.MovementRoutes(router, alertMovementController, receiveMovementController)

    alertDoorService, receiveDoorService, rabbitMQDoor := infresDoor.InitDoor()
    alertDoorController := controllersDoor.NewAlertDoorController(alertDoorService)
    receiveDoorController := controllersDoor.NewReceiveDoorController(receiveDoorService)
    routesDoor.DoorRoutes(router, alertDoorController, receiveDoorController)

    alertWindowService, receiveWindowService, rabbitMQWindow := infresWindow.InitWindow()
    alertWindowController := controllerWindow.NewAlertWindowController(alertWindowService)
    receiveWindowController := controllerWindow.NewReceiveWindowController(receiveWindowService)
    routesWindow.WindowRoutes(router, alertWindowController, receiveWindowController)

    alertGasService, receiveGasService, rabbitMQGas := infresGas.InitGas()
    alertGasController := controllerGas.NewAlertGasController(alertGasService)
    receiveGasController := controllerGas.NewReceiveGasController(receiveGasService)
    routesGas.GasRoutes(router, alertGasController, receiveGasController)

    defer rabbitMQWeather.Close()
    defer rabbitMQGas.Close()
    defer rabbitMQMovement.Close()
    defer rabbitMQLight.Close()
    defer rabbitMQDoor.Close()
    defer rabbitMQWindow.Close()

    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}