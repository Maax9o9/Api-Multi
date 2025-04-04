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
		log.Printf("Error al declarar la cola: %v", err)
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
		log.Printf("Error al enlazar la cola: %v", err)
		return nil, err
	}

	return &RabbitConsumer{
		Connection: conn,
		Channel:    ch,
		QueueName:  q.Name,
	}, nil
}

func (r *RabbitConsumer) ConsumeMessages(processMessage func(body []byte)) error {
	log.Printf("Iniciando consumo de mensajes de la cola: %s", r.QueueName)

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
		log.Printf("Error al consumir mensajes de RabbitMQ: %v", err)
		return err
	}

	log.Printf("Consumidor configurado correctamente. Esperando mensajes...")

	// Canal para notificar cuando la goroutine termina (para debugging)
	done := make(chan bool)

	go func() {
		messageCount := 0
		for d := range msgs {
			messageCount++
			log.Printf("Mensaje #%d recibido ", messageCount)
			log.Printf("Contenido: %s", d.Body)
			log.Printf("Exchange: %s, RoutingKey: %s", d.Exchange, d.RoutingKey)

			processMessage(d.Body)
			log.Printf("Mensaje procesado correctamente")
		}
		done <- true
	}()

	// Log adicional para confirmar que la goroutine está en ejecución
	log.Println("Goroutine de consumo iniciada. Esperando mensajes...")

	// Mantén este método ejecutándose
	<-done
	return nil
}

func (r *RabbitConsumer) Close() {
	if err := r.Channel.Close(); err != nil {
		log.Printf("Error al cerrar el canal de RabbitMQ: %v", err)
	}
	if err := r.Connection.Close(); err != nil {
		log.Printf("Error al cerrar la conexión de RabbitMQ: %v", err)
	}
}
