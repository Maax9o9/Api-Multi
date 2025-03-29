package routes

import (
    "Multi/src/incidencies/infrestructure/controllers"
    "github.com/gin-gonic/gin"
)

func IncidenciesRoutes(router *gin.Engine, showController *controllers.ShowIncidenciesController, incrementController *controllers.IncrementIncidenciesController) {
    router.GET("/api/incidencies", showController.GetAllIncidencies)
    router.GET("/api/incidencies/:type", showController.GetIncidencyByType)
    router.POST("/api/incidencies/:type/increment", incrementController.IncrementIncidency)
}