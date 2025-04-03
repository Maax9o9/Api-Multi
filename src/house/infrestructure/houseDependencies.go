package infrestructure

import (
    "Multi/src/house/application"
    "Multi/src/house/infrestructure/controllers"
)

func InitHouse() (*controllers.AddHouseController, *controllers.ShowHouseController, *controllers.EditHouseController) {
    houseRepo := NewPostgres()

    imageHandlerUseCase := application.NewImageHandlerUseCase()
    locationHandlerUseCase := application.NewLocationHandlerUseCase(houseRepo)

    addHouseController := controllers.NewAddHouseController(imageHandlerUseCase, locationHandlerUseCase)

    showHouseUseCase := application.NewShowHouseUseCase(houseRepo)
    editHouseUseCase := application.NewEditHouseUseCase(houseRepo)

    showHouseController := controllers.NewShowHouseController(showHouseUseCase)
    editHouseController := controllers.NewEditHouseController(editHouseUseCase)

    return addHouseController, showHouseController, editHouseController
}