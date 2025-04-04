package application

import "Multi/src/house/domain"

type UpdateHouseImageUseCase struct {
	houseRepository domain.HouseRepository
}

func NewUpdateHouseImageUseCase(houseRepository domain.HouseRepository) *UpdateHouseImageUseCase {
	return &UpdateHouseImageUseCase{
		houseRepository: houseRepository,
	}
}

func (uc *UpdateHouseImageUseCase) Execute(houseID int, imagePath string) error {
	// Obtener la casa actual
	house, err := uc.houseRepository.GetByID(houseID)
	if err != nil {
		return err
	}

	// Actualizar la ruta de la imagen
	house.Image = imagePath

	// Guardar los cambios
	return uc.houseRepository.Update(house)
}
