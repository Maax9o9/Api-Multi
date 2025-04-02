package main

import (
    infresHouse "Multi/src/house/infrestructure"
    routesHouse "Multi/src/house/infrestructure/routes"
    infresNotification "Multi/src/notification/infrestructure"
    routesNotification "Multi/src/notification/infrestructure/routes"
    infresUser "Multi/src/user/infrestructure"
    routesUser "Multi/src/user/infrestructure/routes"
    infresWeather "Multi/src/weather"
    weatherControllers "Multi/src/weather/infrestructure/controllers"
    routesWeather "Multi/src/weather/infrestructure/routes"
    infresLight "Multi/src/interruptors/light/infrestructure"
    routesLight "Multi/src/interruptors/light/infrestructure/routes"
    lightControllers "Multi/src/interruptors/light/infrestructure/controllers"
    infresMovement "Multi/src/interruptors/movement/infrestructure"
    routesMovement "Multi/src/interruptors/movement/infrestructure/routes"
    controllerMovement "Multi/src/interruptors/movement/infrestructure/controllers"
    infresDoor "Multi/src/interruptors/door/infrestructure"
    routesDoor "Multi/src/interruptors/door/infrestructure/routes"
    controllerDoor "Multi/src/interruptors/door/infrestructure/controllers"
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

    weatherService, rabbitMQWeather := infresWeather.InitWeather()
    weatherController := weatherControllers.NewWeatherController(weatherService)
    routesWeather.WeatherRoutes(router, weatherController)

    createUserController, showUserController, loginController := infresUser.InitUser()
    routesUser.UserRoutes(router, createUserController, showUserController, loginController)

    addHouseController, showHouseController, editHouseController := infresHouse.InitHouse()
    routesHouse.HouseRoutes(router, addHouseController, showHouseController, editHouseController)

    createNotificationController, showNotificationController := infresNotification.InitNotification()
    routesNotification.NotificationRoutes(router, createNotificationController, showNotificationController)

    lightService, rabbitMQLight := infresLight.InitLight()
    lightController := lightControllers.NewLightController(lightService)
    routesLight.LightRoutes(router, lightController)

    movementService, rabbitMQMovement := infresMovement.InitMovement()
    movementController := controllerMovement.NewMovementController(movementService)
    routesMovement.MovementRoutes(router, movementController)

    doorService, rabbitMQDoor := infresDoor.InitDoor()
    doorController := controllerDoor.NewDoorController(doorService)
    routesDoor.DoorRoutes(router, doorController)

    windowService, rabbitMQWindow := infresWindow.InitWindow()
    windowController := controllerWindow.NewWindowController(windowService)
    routesWindow.WindowRoutes(router, windowController)

    gasService, rabbitMQGas := infresGas.InitGas()
    gasController := controllerGas.NewGasController(gasService)
    routesGas.GasRoutes(router, gasController)

    defer rabbitMQWeather.Close()
    defer rabbitMQGas.Close()
    defer rabbitMQMovement.Close()
    defer rabbitMQLight.Close()
    defer rabbitMQMovement.Close()
    defer rabbitMQDoor.Close()
    defer rabbitMQWindow.Close()
    defer rabbitMQGas.Close()

    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}