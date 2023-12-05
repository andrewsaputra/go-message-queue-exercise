package main

import (
	"andrewsaputra/go-message-queue-exercise/apiserver/api"
	"andrewsaputra/go-message-queue-exercise/apiserver/impl"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var startTime time.Time = time.Now()

func main() {
	router := gin.Default()
	router.GET("/status", StatusCheck)

	var apiService api.ApiService
	var userAccessor api.UserDataAccessor
	var productAccessor api.ProductDataAccessor
	var mqService api.MQService

	userAccessor, err := impl.ConstructUserPGAccessor("postgres://demouser:password@localhost:5432/demo_message_queue")
	if err != nil {
		log.Panic(err)
	}
	defer userAccessor.Close()

	productAccessor, err = impl.ConstructProductPGAccessor("postgres://demouser:password@localhost:5432/demo_message_queue")
	if err != nil {
		log.Panic(err)
	}
	defer productAccessor.Close()

	mqService, err = impl.ConstructRabbitMQService("amqp://guest:guest@127.0.0.1:5672")
	if err != nil {
		log.Panic(err)
	}
	defer mqService.Close()

	apiService = impl.ConstructApiServiceImpl(userAccessor, productAccessor, mqService)
	apiHandler := impl.ConstructApiHandler(apiService)
	router.GET("/users/:id", apiHandler.GetUser)
	router.DELETE("/users/:id", apiHandler.DeleteUser)
	router.POST("/users", apiHandler.AddUser)
	router.GET("/products/:id", apiHandler.GetProduct)
	router.DELETE("/products/:id", apiHandler.DeleteProduct)
	router.POST("/products", apiHandler.AddProduct)

	router.Run(":3000")
}

func StatusCheck(c *gin.Context) {
	body := make(map[string]any)
	body["status"] = "Healthy"
	body["started_at"] = startTime.Format(time.RFC822Z)

	c.JSON(http.StatusOK, body)
}
