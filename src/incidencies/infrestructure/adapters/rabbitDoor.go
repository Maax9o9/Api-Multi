package adapters

import (
    "Multi/src/incidencies/domain/entities"
    "log"
)

type RabbitDoorAdapter struct {
    *RabbitAdapter
}

func NewRabbitDoorAdapter(rabbitURL string) (*RabbitDoorAdapter, error) {
    adapter, err := NewRabbitAdapter(rabbitURL, "door_opened_queue")
    if err != nil {
        return nil, err
    }
    return &RabbitDoorAdapter{adapter}, nil
}

func (rda *RabbitDoorAdapter) PublishDoorIncidency(incidency entities.Incidency) error {
    log.Println("Publicando incidencia de puerta abierta...")
    return rda.PublishIncidency(incidency)
}