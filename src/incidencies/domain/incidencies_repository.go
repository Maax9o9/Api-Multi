package domain

import "Multi/src/incidencies/domain/entities"

type IncidenciesRepository interface {
    Increment(typeNotification string) (*entities.Incidency, error)
    GetAll() ([]entities.Incidency, error)
    GetByType(typeNotification string) (*entities.Incidency, error)
}