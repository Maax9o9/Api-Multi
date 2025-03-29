package adapters

import (
    "Multi/src/incidencies/domain/entities"
    "log"
)

type RabbitWindowAdapter struct {
    *RabbitAdapter
}

func NewRabbitWindowAdapter(rabbitURL string) (*RabbitWindowAdapter, error) {
    adapter, err := NewRabbitAdapter(rabbitURL, "window_opened_queue")
    if err != nil {
        return nil, err
    }
    return &RabbitWindowAdapter{adapter}, nil
}

func (rwa *RabbitWindowAdapter) PublishWindowIncidency(incidency entities.Incidency) error {
    log.Println("Publicando incidencia de ventana abierta...")
    return rwa.PublishIncidency(incidency)
}