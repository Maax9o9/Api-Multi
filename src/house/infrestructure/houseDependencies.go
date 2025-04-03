package infrestructure

import (
    "Multi/src/house/application"
    "Multi/src/house/infrestructure/controllers"
)

func InitHouse() (*controllers.AddHouseController, *controllers.ShowAllHousesController, *controllers.ShowHouseByIDController, *controllers.EditHouseController) {
    houseRepo := NewPostgres()

    imageHandlerUseCase := application.NewImageHandlerUseCase()
    locationHandlerUseCase := application.NewLocationHandlerUseCase(houseRepo)
    showAllHousesUseCase := application.NewShowAllHousesUseCase(houseRepo)
    showHouseByIDUseCase := application.NewShowHouseByIDUseCase(houseRepo)
    editHouseUseCase := application.NewEditHouseUseCase(houseRepo)

    addHouseController := controllers.NewAddHouseController(imageHandlerUseCase, locationHandlerUseCase)
    showAllHousesController := controllers.NewShowAllHousesController(showAllHousesUseCase)
    showHouseByIDController := controllers.NewShowHouseByIDController(showHouseByIDUseCase)
    editHouseController := controllers.NewEditHouseController(editHouseUseCase)

    return addHouseController, showAllHousesController, showHouseByIDController, editHouseController
}