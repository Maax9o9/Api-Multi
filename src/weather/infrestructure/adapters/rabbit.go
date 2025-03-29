package adapters

import (
    "log"

    amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitConsumer struct {
    Connection *amqp.Connection
    Channel    *amqp.Channel
    QueueName  string
}

func NewRabbitConsumer(rabbitURL, exchangeName, queueName, routingKey string) (*RabbitConsumer, error) {
    conn, err := amqp.Dial(rabbitURL)
    if err != nil {
        return nil, err
    }

    ch, err := conn.Channel()
    if err != nil {
        conn.Close()
        return nil, err
    }

    err = ch.ExchangeDeclare(
        exchangeName, // name
        "direct",     // type
        true,         // durable
        false,        // auto-deleted
        false,        // internal
        false,        // no-wait
        nil,          // arguments
    )
    if err != nil {
        ch.Close()
        conn.Close()
        return nil, err
    }

    q, err := ch.QueueDeclare(
        queueName, // name
        true,      // durable
        false,     // delete when unused
        false,     // exclusive
        false,     // no-wait
        nil,       // arguments
    )
    if err != nil {
        ch.Close()
        conn.Close()
        return nil, err
    }

    err = ch.QueueBind(
        q.Name,       // queue name
        routingKey,   // routing key
        exchangeName, // exchange
        false,        // no-wait
        nil,          // arguments
    )
    if err != nil {
        ch.Close()
        conn.Close()
        return nil, err
    }

    return &RabbitConsumer{
        Connection: conn,
        Channel:    ch,
        QueueName:  q.Name,
    }, nil
}

func (r *RabbitConsumer) ConsumeMessages(processMessage func(body []byte)) error {
    msgs, err := r.Channel.Consume(
        r.QueueName, // queue
        "",          // consumer
        true,        // auto-ack
        false,       // exclusive
        false,       // no-local
        false,       // no-wait
        nil,         // args
    )
    if err != nil {
        return err
    }

    go func() {
        for d := range msgs {
            log.Printf("Received a message: %s", d.Body)
            processMessage(d.Body)
        }
    }()

    return nil
}

func (r *RabbitConsumer) Close() {
    r.Channel.Close()
    r.Connection.Close()
}