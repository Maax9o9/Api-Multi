package application

import (
    "Multi/src/notification/domain"
    "Multi/src/notification/domain/entities"
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

func (uc *CreateNotificationUseCase) CreateNotification(houseID int, message string, typeNotification string) (*entities.Notification, error) {
    notification := &entities.Notification{
        HouseID:          houseID,
        Date:             time.Now(),
        Message:          message,
        TypeNotification: typeNotification,
    }

    err := uc.repo.Create(notification)
    if err != nil {
        return nil, err
    }

    return notification, nil
}