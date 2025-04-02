package adapters

import (
    "github.com/streadway/amqp"
    "log"
)

type RabbitConsumer struct {
    connection *amqp.Connection
    channel    *amqp.Channel
    queue      amqp.Queue
}

func NewRabbitConsumer(connectionString, queueName string) (*RabbitConsumer, error) {
    conn, err := amqp.Dial(connectionString)
    if err != nil {
        log.Printf("Error al conectar con RabbitMQ: %v", err)
        return nil, err
    }

    channel, err := conn.Channel()
    if err != nil {
        log.Printf("Error al crear el canal de RabbitMQ: %v", err)
        return nil, err
    }

    queue, err := channel.QueueDeclare(
        queueName, // Nombre de la cola
        true,      // Durable
        false,     // Auto-delete
        false,     // Exclusive
        false,     // No-wait
        nil,       // Arguments
    )
    if err != nil {
        log.Printf("Error al declarar la cola de RabbitMQ: %v", err)
        return nil, err
    }

    return &RabbitConsumer{
        connection: conn,
        channel:    channel,
        queue:      queue,
    }, nil
}

func (r *RabbitConsumer) ConsumeMessages(processMessage func(body []byte)) error {
    msgs, err := r.channel.Consume(
        r.queue.Name, // Nombre de la cola
        "",           // Consumer
        true,         // Auto-ack
        false,        // Exclusive
        false,        // No-local
        false,        // No-wait
        nil,          // Args
    )
    if err != nil {
        log.Printf("Error al consumir mensajes de RabbitMQ: %v", err)
        return err
    }

    for msg := range msgs {
        processMessage(msg.Body)
    }

    return nil
}

func (r *RabbitConsumer) Close() {
    if err := r.channel.Close(); err != nil {
        log.Printf("Error al cerrar el canal de RabbitMQ: %v", err)
    }
    if err := r.connection.Close(); err != nil {
        log.Printf("Error al cerrar la conexión de RabbitMQ: %v", err)
    }
}