package adapters

import (
	"context"
	"log"
	"time"

	"github.com/streadway/amqp"
)

type RabbitPublisher struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Exchange   string
}

// NewRabbitPublisher crea una nueva instancia de publicador RabbitMQ
func NewRabbitPublisher(rabbitURL, exchangeName string) (*RabbitPublisher, error) {
	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		log.Printf("Error al conectar con RabbitMQ: %v", err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		log.Printf("Error al crear el canal de RabbitMQ: %v", err)
		return nil, err
	}

	err = ch.ExchangeDeclare(
		exchangeName, // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		ch.Close()
		conn.Close()
		log.Printf("Error al declarar el exchange: %v", err)
		return nil, err
	}

	return &RabbitPublisher{
		Connection: conn,
		Channel:    ch,
		Exchange:   exchangeName,
	}, nil
}

// PublishMessage publica un mensaje en RabbitMQ
func (r *RabbitPublisher) PublishMessage(routingKey string, message []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.Channel.Publish(
		r.Exchange, // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)
	if err != nil {
		log.Printf("Error al publicar mensaje: %v", err)
		return err
	}

	log.Printf("Mensaje publicado exitosamente en la ruta %s: %s", routingKey, string(message))
	ctx.Done()
	return nil
}

// Close cierra las conexiones con RabbitMQ
func (r *RabbitPublisher) Close() {
	if r.Channel != nil {
		r.Channel.Close()
	}
	if r.Connection != nil {
		r.Connection.Close()
	}
}
