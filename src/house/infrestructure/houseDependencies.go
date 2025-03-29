package infrestructure

import (
    "Multi/src/house/application"
    "Multi/src/house/infrestructure/controllers"
)

func InitHouse() (*controllers.AddHouseController, *controllers.ShowHouseController, *controllers.EditHouseController) {
    houseRepo := NewPostgres()

    addHouseUseCase := application.NewAddHouseUseCase(houseRepo)
    showHouseUseCase := application.NewShowHouseUseCase(houseRepo)
    editHouseUseCase := application.NewEditHouseUseCase(houseRepo)

    addHouseController := controllers.NewAddHouseController(addHouseUseCase)
    showHouseController := controllers.NewShowHouseController(showHouseUseCase)
    editHouseController := controllers.NewEditHouseController(editHouseUseCase)

    return addHouseController, showHouseController, editHouseController
}