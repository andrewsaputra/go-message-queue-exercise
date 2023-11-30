package impl

import (
	"andrewsaputra/go-message-queue-exercise/apiserver/api"
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ConstructRabbitMQService() (api.MQService, error) {
	service := &RabbitMQService{}
	if err := service.connect("amqp://guest:guest@127.0.0.1:5672"); err != nil {
		return nil, err
	}

	return service, nil
}

type RabbitMQService struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func (this *RabbitMQService) Publish(payload api.MQPublish) error {
	queue, err := this.Channel.QueueDeclare(
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

	err = this.Channel.PublishWithContext(
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
	if this.Channel != nil {
		this.Channel.Close()
	}

	if this.Connection != nil {
		this.Connection.Close()
	}
}

func (this *RabbitMQService) connect(url string) error {
	conn, err := amqp.Dial(url)
	if err != nil {
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	this.Connection = conn
	this.Channel = ch
	return nil
}
