package application

import (
    "Multi/src/notification/domain"
    "Multi/src/notification/domain/entities"
)

type ShowNotificationByIDUseCase struct {
    repo domain.NotificationRepository
}

func NewShowNotificationByIDUseCase(repo domain.NotificationRepository) *ShowNotificationByIDUseCase {
    return &ShowNotificationByIDUseCase{
        repo: repo,
    }
}

func (uc *ShowNotificationByIDUseCase) GetNotificationByID(id int) (*entities.Notification, error) {
    return uc.repo.GetByID(id)
}