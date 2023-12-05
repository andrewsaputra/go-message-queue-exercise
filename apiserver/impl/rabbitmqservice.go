package impl

import (
	"andrewsaputra/go-message-queue-exercise/apiserver/api"
	"context"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ConstructRabbitMQService(addr string) (api.MQService, error) {
	service := &RabbitMQService{Address: addr}
	if err := service.reconnect(); err != nil {
		return nil, err
	}

	return service, nil
}

type RabbitMQService struct {
	Address    string
	Connection *amqp.Connection
	Mutex      sync.RWMutex
}

func (this *RabbitMQService) Publish(payload api.MQPublish) error {
	this.Mutex.RLock()
	isClosed := this.Connection.IsClosed()
	this.Mutex.RUnlock()

	if isClosed {
		if err := this.reconnect(); err != nil {
			return err
		}
	}

	channel, err := this.Connection.Channel()
	if err != nil {
		return err
	}

	queue, err := channel.QueueDeclare(
		payload.Queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = channel.PublishWithContext(
		ctx,
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(payload.Message),
		},
	)

	return err
}

func (this *RabbitMQService) Close() {
	if this.Connection != nil && !this.Connection.IsClosed() {
		this.Connection.Close()
	}
}

func (this *RabbitMQService) reconnect() error {
	this.Mutex.Lock()
	defer this.Mutex.Unlock()

	conn, err := amqp.Dial(this.Address)
	if err != nil {
		return err
	}

	this.Connection = conn
	return nil
}
