package main

import (
	"andrewsaputra/go-message-queue-exercise/apiserver/api"
	"andrewsaputra/go-message-queue-exercise/apiserver/impl"
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

	apiService = impl.ConstructApiServiceImpl(userAccessor, productAccessor, mqService)
	apiHandler := impl.ConstructApiHandler(apiService)
	router.GET("/user/:id", apiHandler.GetUser)
	router.POST("/user", apiHandler.AddUser)
	router.GET("/product/:id", apiHandler.GetProduct)
	router.POST("/product", apiHandler.AddProduct)

	router.Run(":3000")
}

func StatusCheck(c *gin.Context) {
	body := make(map[string]any)
	body["status"] = "Healthy"
	body["started_at"] = startTime.Format(time.RFC822Z)

	c.JSON(http.StatusOK, body)
}
