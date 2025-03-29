package application

import (
    "Multi/src/notification/domain"
    "Multi/src/notification/domain/entities"
)

type ShowNotificationUseCase struct {
    repo domain.NotificationRepository
}

func NewShowNotificationUseCase(repo domain.NotificationRepository) *ShowNotificationUseCase {
    return &ShowNotificationUseCase{
        repo: repo,
    }
}

func (uc *ShowNotificationUseCase) GetAllNotifications() ([]entities.Notification, error) {
    return uc.repo.GetAll()
}

func (uc *ShowNotificationUseCase) GetNotificationByID(id int) (*entities.Notification, error) {
    return uc.repo.GetByID(id)
}