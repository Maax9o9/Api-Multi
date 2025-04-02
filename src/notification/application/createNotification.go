package application

import (
    "Multi/src/notification/domain"
    "Multi/src/notification/domain/entities"
    "errors"
    "time"
)

type CreateNotificationUseCase struct {
    repo domain.NotificationRepository
}

func NewCreateNotificationUseCase(repo domain.NotificationRepository) *CreateNotificationUseCase {
    return &CreateNotificationUseCase{
        repo: repo,
    }
}

func (uc *CreateNotificationUseCase) CreateNotification(houseID int, sensorID int, sensorType string, message string) (*entities.Notification, error) {
    validSensorTypes := map[string]bool{
        "GasSensor":     true,
        "MotionSensor":  true,
        "DoorSensor":    true,
        "WindowSensor":  true,
        "LedControl":    true,
    }
    if !validSensorTypes[sensorType] {
        return nil, errors.New("invalid sensorType: must be one of 'GasSensor', 'MotionSensor', 'DoorSensor', 'WindowSensor', or 'LedControl'")
    }

    notification := &entities.Notification{
        HouseID:    houseID,
        SensorID:   sensorID,
        SensorType: sensorType,
        Date:       time.Now(),
        Message:    message,
    }

    err := uc.repo.Create(notification)
    if err != nil {
        return nil, err
    }

    return notification, nil
}