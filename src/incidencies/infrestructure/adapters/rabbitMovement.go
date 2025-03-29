package adapters

import (
    "Multi/src/incidencies/domain/entities"
    "log"
)

type RabbitMovementAdapter struct {
    *RabbitAdapter
}

func NewRabbitMovementAdapter(rabbitURL string) (*RabbitMovementAdapter, error) {
    adapter, err := NewRabbitAdapter(rabbitURL, "movement_detected_queue")
    if err != nil {
        return nil, err
    }
    return &RabbitMovementAdapter{adapter}, nil
}

func (rma *RabbitMovementAdapter) PublishMovementIncidency(incidency entities.Incidency) error {
    log.Println("Publicando incidencia de movimiento detectado...")
    return rma.PublishIncidency(incidency)
}