package impl

import (
	"andrewsaputra/go-message-queue-exercise/consumer/api"
	"fmt"

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
}

func (this *RabbitMQService) Subscribe(queueName string, consumer api.ConsumerService) error {
	channel, err := this.Connection.Channel()
	if err != nil {
		return err
	}

	queue, err := channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	msgs, err := channel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	go func() {
		for d := range msgs {
			if err := consumer.OnConsumed(d.Body); err != nil {
				fmt.Println("OnConsumed error : ", err)
				continue
			}

			d.Ack(false)
		}
	}()

	return err
}

func (this *RabbitMQService) Close() {
	if this.Connection != nil && !this.Connection.IsClosed() {
		this.Connection.Close()
	}
}

func (this *RabbitMQService) reconnect() error {
	conn, err := amqp.Dial(this.Address)
	if err != nil {
		return err
	}

	this.Connection = conn
	return nil
}
