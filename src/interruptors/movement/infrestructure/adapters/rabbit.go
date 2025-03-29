package adapters

import (
    "Multi/src/interruptors/movement/domain/entities"
    "encoding/json"
    "log"

    "github.com/streadway/amqp"
)

type RabbitAdapter struct {
    channel *amqp.Channel
    queue   string
}

func NewRabbitAdapter(channel *amqp.Channel, queue string) *RabbitAdapter {
    return &RabbitAdapter{
        channel: channel,
        queue:   queue,
    }
}

func (r *RabbitAdapter) PublishMovementCommand(command entities.MovementCommand) error {
    message, err := json.Marshal(command)
    if err != nil {
        log.Printf("Error al serializar el comando de movimiento: %v", err)
        return err
    }

    err = r.channel.Publish(
        "",         // exchange
        r.queue,    // routing key
        false,      // mandatory
        false,      // immediate
        amqp.Publishing{
            ContentType: "application/json",
            Body:        message,
        },
    )
    if err != nil {
        log.Printf("Error al publicar el comando de movimiento en RabbitMQ: %v", err)
        return err
    }

    log.Printf("Comando de movimiento publicado en RabbitMQ: %+v", command)
    return nil
}