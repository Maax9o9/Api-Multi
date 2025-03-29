package adapters

import (
    "encoding/json"
    "log"

    amqp "github.com/streadway/amqp"
    "Multi/src/incidencies/domain/entities"
)

type RabbitAdapter struct {
    conn    *amqp.Connection
    channel *amqp.Channel
    queue   string
}

func NewRabbitAdapter(rabbitURL, queueName string) (*RabbitAdapter, error) {
    conn, err := amqp.Dial(rabbitURL)
    if err != nil {
        log.Printf("Error connecting to RabbitMQ: %v", err)
        return nil, err
    }

    channel, err := conn.Channel()
    if err != nil {
        log.Printf("Error opening a channel in RabbitMQ: %v", err)
        return nil, err
    }

    _, err = channel.QueueDeclare(
		queueName, // Queue name
        true,      // Durable
        false,     // Auto-delete
        false,     // Exclusive
        false,     // No-wait
        nil,       // Arguments
    )
    if err != nil {
        log.Printf("Error declaring the queue: %v", err)
        return nil, err
    }

    return &RabbitAdapter{
        conn:    conn,
        channel: channel,
        queue:   queueName,
    }, nil
}

func (ra *RabbitAdapter) ConsumeMessages(processMessage func(body []byte)) error {
    msgs, err := ra.channel.Consume(
        ra.queue, // Queue
        "",       // Consumer
        true,     // Auto-ack
        false,    // Exclusive
        false,    // No-local
        false,    // No-wait
        nil,      // Args
    )
    if err != nil {
        return err
    }

    go func() {
        for d := range msgs {
            log.Printf("Message received: %s", d.Body)
            processMessage(d.Body)
        }
    }()

    return nil
}

func (ra *RabbitAdapter) PublishMessage(message []byte) error {
    err := ra.channel.Publish(
        "",       // Exchange
        ra.queue, // Routing key (queue name)
        false,    // Mandatory
        false,    // Immediate
        amqp.Publishing{
            ContentType: "application/json",
            Body:        message,
        },
    )
    if err != nil {
        log.Printf("Error publishing the message to RabbitMQ: %v", err)
        return err
    }

    log.Printf("Message published to RabbitMQ: %s", message)
    return nil
}

func (ra *RabbitAdapter) PublishIncidency(incidency entities.Incidency) error {
    log.Printf("Publishing incidency: %s", incidency.TypeNotification)

    message, err := json.Marshal(incidency)
    if err != nil {
        log.Printf("Error serializing the incidency: %v", err)
        return err
    }

    return ra.PublishMessage(message)
}

func (ra *RabbitAdapter) Close() {
    if err := ra.channel.Close(); err != nil {
        log.Printf("Error closing RabbitMQ channel: %v", err)
    }
    if err := ra.conn.Close(); err != nil {
        log.Printf("Error closing RabbitMQ connection: %v", err)
    }
}