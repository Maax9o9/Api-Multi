package adapters

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTPublisher struct {
	Client  mqtt.Client
	Options *mqtt.ClientOptions
}

// NewMQTTPublisher crea una nueva instancia de publicador MQTT
func NewMQTTPublisher(brokerURL, clientID string, username, password string) (*MQTTPublisher, error) {
	// Configurar opciones del cliente MQTT
	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerURL)
	opts.SetClientID(clientID)

	// Configurar credenciales si se proporcionan
	if username != "" {
		opts.SetUsername(username)
		opts.SetPassword(password)
	}

	// Configurar callbacks
	opts.SetOnConnectHandler(func(client mqtt.Client) {
		log.Println("Conectado al broker MQTT")
	})

	opts.SetConnectionLostHandler(func(client mqtt.Client, err error) {
		log.Printf("Conexión MQTT perdida: %v", err)
	})

	// Conexión persistente
	opts.SetAutoReconnect(true)
	opts.SetMaxReconnectInterval(1 * time.Minute)
	opts.SetKeepAlive(30 * time.Second)

	// Crear el cliente MQTT
	client := mqtt.NewClient(opts)

	// Conectar al broker
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		return nil, fmt.Errorf("error al conectar con el broker MQTT: %v", token.Error())
	}

	return &MQTTPublisher{
		Client:  client,
		Options: opts,
	}, nil
}

// PublishMessage publica un mensaje en un topic MQTT
func (m *MQTTPublisher) PublishMessage(topic string, message []byte) error {
	// Asegurar que estamos conectados
	if !m.Client.IsConnected() {
		log.Println("Reconectando al broker MQTT...")
		token := m.Client.Connect()
		if token.Wait() && token.Error() != nil {
			return fmt.Errorf("error al reconectar con MQTT: %v", token.Error())
		}
	}

	// Publicar el mensaje
	token := m.Client.Publish(topic, 0, false, message)
	token.Wait()

	if token.Error() != nil {
		return fmt.Errorf("error al publicar mensaje MQTT: %v", token.Error())
	}

	log.Printf("Mensaje MQTT publicado exitosamente en el topic %s: %s", topic, string(message))
	return nil
}

// Close cierra la conexión MQTT
func (m *MQTTPublisher) Close() {
	if m.Client.IsConnected() {
		m.Client.Disconnect(250)
	}
}
