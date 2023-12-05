package impl

import (
	"andrewsaputra/go-message-queue-exercise/apiserver/api"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ConstructApiHandler(apiService api.ApiService) *ApiHandler {
	validator := validator.New(validator.WithRequiredStructEnabled())

	return &ApiHandler{
		Validator: validator,
		Service:   apiService,
	}
}

type ApiHandler struct {
	Validator *validator.Validate
	Service   api.ApiService
}

func (this *ApiHandler) GetUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		this.writeResponse(
			c,
			api.ApiHandlerResponse{Code: http.StatusBadRequest, Error: errors.New("invalid id provided")},
		)
		return
	}

	response := this.Service.GetUser(userId)
	this.writeResponse(c, response)
}

func (this *ApiHandler) DeleteUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		this.writeResponse(
			c,
			api.ApiHandlerResponse{Code: http.StatusBadRequest, Error: errors.New("invalid id provided")},
		)
		return
	}

	response := this.Service.DeleteUser(userId)
	this.writeResponse(c, response)
}

func (this *ApiHandler) AddUser(c *gin.Context) {
	var dto api.AddUserDTO
	if err := c.BindJSON(&dto); err != nil {
		this.writeResponse(
			c,
			api.ApiHandlerResponse{Code: http.StatusBadRequest, Error: err},
		)
		return
	}

	if err := this.Validator.Struct(dto); err != nil {
		this.writeResponse(
			c,
			api.ApiHandlerResponse{Code: http.StatusBadRequest, Error: err},
		)
		return
	}

	response := this.Service.AddUser(dto)
	this.writeResponse(c, response)
}

func (this *ApiHandler) GetProduct(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		this.writeResponse(
			c,
			api.ApiHandlerResponse{Code: http.StatusBadRequest, Error: errors.New("invalid id provided")},
		)
		return
	}

	response := this.Service.GetProduct(productId)
	this.writeResponse(c, response)
}

func (this *ApiHandler) DeleteProduct(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		this.writeResponse(
			c,
			api.ApiHandlerResponse{Code: http.StatusBadRequest, Error: errors.New("invalid id provided")},
		)
		return
	}

	response := this.Service.DeleteProduct(userId)
	this.writeResponse(c, response)
}

func (this *ApiHandler) AddProduct(c *gin.Context) {
	var dto api.AddProductDTO
	if err := c.BindJSON(&dto); err != nil {
		this.writeResponse(
			c,
			api.ApiHandlerResponse{Code: http.StatusBadRequest, Error: err},
		)
		return
	}

	if err := this.Validator.Struct(dto); err != nil {
		this.writeResponse(
			c,
			api.ApiHandlerResponse{Code: http.StatusBadRequest, Error: err},
		)
		return
	}

	response := this.Service.AddProduct(dto)
	this.writeResponse(c, response)
}

func (this *ApiHandler) writeResponse(c *gin.Context, response api.ApiHandlerResponse) {
	if response.Error != nil {
		c.JSON(response.Code, gin.H{"error": response.Error.Error()})
		return
	}

	c.JSON(response.Code, response.Body)
}
