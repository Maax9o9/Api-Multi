package adapters

import (
    "Multi/src/incidencies/domain/entities"
    "log"
)

type RabbitGasAdapter struct {
    *RabbitAdapter
}

func NewRabbitGasAdapter(rabbitURL string) (*RabbitGasAdapter, error) {
    adapter, err := NewRabbitAdapter(rabbitURL, "gas_detected_queue")
    if err != nil {
        return nil, err
    }
    return &RabbitGasAdapter{adapter}, nil
}

func (rga *RabbitGasAdapter) PublishGasIncidency(incidency entities.Incidency) error {
    log.Println("Publicando incidencia de gas detectado...")
    return rga.PublishIncidency(incidency)
}