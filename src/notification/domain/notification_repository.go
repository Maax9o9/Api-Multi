package domain

import "Multi/src/notification/domain/entities"

type NotificationRepository interface {
    Create(notification *entities.Notification) error
    GetAll() ([]entities.Notification, error)
    GetByID(id int) (*entities.Notification, error)
}