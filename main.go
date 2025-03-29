package main

import (
    infresHouse "Multi/src/house/infrestructure"
    routesHouse "Multi/src/house/infrestructure/routes"
    infresIncidencies "Multi/src/incidencies/infrestructure"
    incidenciesControllers "Multi/src/incidencies/infrestructure/controllers"
    routesIncidencies "Multi/src/incidencies/infrestructure/routes"
    infresNotification "Multi/src/notification/infrestructure"
    routesNotification "Multi/src/notification/infrestructure/routes"
    infresUser "Multi/src/user/infrestructure"
    routesUser "Multi/src/user/infrestructure/routes"
    infresWeather "Multi/src/weather"
    weatherControllers "Multi/src/weather/infrestructure/controllers"
    routesWeather "Multi/src/weather/infrestructure/routes"
    infresLight "Multi/src/interruptors/light/infrestructure"
    routesLight "Multi/src/interruptors/light/infrestructure/routes"
    infresMovement "Multi/src/interruptors/movement/infrestructure"
    routesMovement "Multi/src/interruptors/movement/infrestructure/routes"
    infresDoor "Multi/src/interruptors/door/infrestructure"
    routesDoor "Multi/src/interruptors/door/infrestructure/routes"
    infresWindow "Multi/src/interruptors/window/infrestructure"
    routesWindow "Multi/src/interruptors/window/infrestructure/routes"
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

    getIncidenciesUseCase, incrementIncidencyUseCase, rabbitGas, rabbitMovement, rabbitDoor, rabbitWindow, _, rabbitMQIncidencies := infresIncidencies.InitIncidencies()
    showIncidenciesController := incidenciesControllers.NewShowIncidenciesController(getIncidenciesUseCase)
    incrementIncidenciesController := incidenciesControllers.NewIncrementIncidenciesController(incrementIncidencyUseCase)
    routesIncidencies.IncidenciesRoutes(router, showIncidenciesController, incrementIncidenciesController)

    createNotificationController, showNotificationController := infresNotification.InitNotification()
    routesNotification.NotificationRoutes(router, createNotificationController, showNotificationController)

    lightController := infresLight.InitLight()
    routesLight.LightRoutes(router, lightController)

    movementController := infresMovement.InitMovement()
    routesMovement.MovementRoutes(router, movementController)

    doorController := infresDoor.InitDoor()
    routesDoor.DoorRoutes(router, doorController)

    windowController := infresWindow.InitWindow()
    routesWindow.WindowRoutes(router, windowController)

    defer rabbitMQWeather.Close()
    defer rabbitMQIncidencies.Close()
    defer rabbitGas.Close()
    defer rabbitMovement.Close()
    defer rabbitDoor.Close()
    defer rabbitWindow.Close()

    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}