package infrestructure

import (
	"Multi/src/core"
	"Multi/src/house/application"
	"Multi/src/house/infrestructure/adapters"
	"Multi/src/house/infrestructure/controllers"
	"Multi/src/house/infrestructure/repositories"
)

// HouseControllers contiene todos los controllers relacionados con casas
type HouseControllers struct {
	CreateHouseController      *controllers.CreateHouseController
	GetHouseByIDController     *controllers.GetHouseByIDController
	GetUserHousesController    *controllers.GetUserHousesController
	UpdateHouseController      *controllers.UpdateHouseController
	DeleteHouseController      *controllers.DeleteHouseController
	UpdateLocationController   *controllers.UpdateLocationController
	UploadImageController      *controllers.UploadImageController
	AddMemberToHouseController *controllers.AddMemberToHouseController
}

// InitializeDependencies configura todas las dependencias necesarias para el módulo de casas
func InitializeDependencies() *HouseControllers {
	// Obtener la conexión a la base de datos usando el Core
	db := core.GetDBPool()

	// Crear adaptadores (implementaciones de las interfaces)
	fileStorage := adapters.NewLocalFileStorage("./uploads/images")

	// Crear repositorios
	houseRepo := repositories.NewPostgresHouseRepository(db)
	memberRepo := repositories.NewPostgresHouseMemberRepository(db)

	// Crear casos de uso para manejo de archivos e imágenes
	imageHandlerUseCase := application.NewImageHandlerUseCase(fileStorage)
	locationHandlerUseCase := application.NewLocationHandlerUseCase()

	// Crear casos de uso para operaciones de casas
	createHouseUseCase := application.NewCreateHouseUseCase(houseRepo, memberRepo)
	getHouseByIDUseCase := application.NewGetHouseByIDUseCase(houseRepo)
	getHousesByUserIDUseCase := application.NewGetHousesByUserIDUseCase(houseRepo)
	updateHouseUseCase := application.NewUpdateHouseUseCase(houseRepo)
	deleteHouseUseCase := application.NewDeleteHouseUseCase(houseRepo)
	updateImageHouse := application.NewUpdateHouseImageUseCase(houseRepo)
	addMemberToHouseUseCase := application.NewAddMemberToHouseUseCase(houseRepo, memberRepo)

	// Crear controladores
	createHouseController := controllers.NewCreateHouseController(
		createHouseUseCase,
		imageHandlerUseCase,
		locationHandlerUseCase,
	)

	getHouseByIDController := controllers.NewGetHouseByIDController(
		getHouseByIDUseCase,
	)

	getUserHousesController := controllers.NewGetUserHousesController(
		getHousesByUserIDUseCase,
	)

	updateHouseController := controllers.NewUpdateHouseController(
		updateHouseUseCase,
		getHouseByIDUseCase,
		imageHandlerUseCase,
		locationHandlerUseCase,
	)

	deleteHouseController := controllers.NewDeleteHouseController(
		deleteHouseUseCase,
		getHouseByIDUseCase,
	)

	updateLocationController := controllers.NewUpdateLocationController(
		locationHandlerUseCase,
		getHouseByIDUseCase,
		updateHouseUseCase,
	)

	uploadImageController := controllers.NewUploadImageController(
		imageHandlerUseCase,
		updateImageHouse,
	)

	addMemberToHouseController := controllers.NewAddMemberToHouseController(
		addMemberToHouseUseCase,
	)

	// Retornar todos los controllers
	return &HouseControllers{
		CreateHouseController:      createHouseController,
		GetHouseByIDController:     getHouseByIDController,
		GetUserHousesController:    getUserHousesController,
		UpdateHouseController:      updateHouseController,
		DeleteHouseController:      deleteHouseController,
		UpdateLocationController:   updateLocationController,
		UploadImageController:      uploadImageController,
		AddMemberToHouseController: addMemberToHouseController,
	}
}
