package main

import (
	"andrewsaputra/go-message-queue-exercise/consumer/api"
	"andrewsaputra/go-message-queue-exercise/consumer/impl"
	"log"
)

func main() {
	var mqService api.MQService
	mqService, err := impl.ConstructRabbitMQService("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Panic(err)
	}
	defer mqService.Close()

	var productAccessor api.ProductDataAccessor
	productAccessor, err = impl.ConstructProductPGAccessor("postgres://demouser:password@localhost:5432/demo_message_queue")
	if err != nil {
		log.Panic(err)
	}

	var consumerService api.ConsumerService
	consumerService = impl.ConstructProductConsumerService(
		"../images",
		"../images/compressed",
		200,
		productAccessor,
	)

	mqService.Subscribe("AddProduct", consumerService)

	var forever chan int
	<-forever
}
