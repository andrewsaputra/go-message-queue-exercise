package impl

import (
	"andrewsaputra/go-message-queue-exercise/apiserver/api"
	"errors"
	"net/http"
	"strconv"
)

func ConstructApiServiceImpl(userAccessor api.UserDataAccessor, productAccessor api.ProductDataAccessor, mqService api.MQService) api.ApiService {
	return &ApiServiceImpl{
		UserAccessor:    userAccessor,
		ProductAccessor: productAccessor,
		MQService:       mqService,
	}
}

type ApiServiceImpl struct {
	UserAccessor    api.UserDataAccessor
	ProductAccessor api.ProductDataAccessor
	MQService       api.MQService
}

func (this *ApiServiceImpl) GetUser(id int) api.ApiHandlerResponse {
	user, err := this.UserAccessor.Get(id)
	if err != nil {
		return api.ApiHandlerResponse{Code: http.StatusInternalServerError, Error: err}
	}
	if user == nil {
		return api.ApiHandlerResponse{Code: http.StatusNotFound, Error: errors.New("data not found")}
	}

	return api.ApiHandlerResponse{Code: http.StatusOK, Body: api.ResponseBody{Data: user}}
}

func (this *ApiServiceImpl) DeleteUser(id int) api.ApiHandlerResponse {
	res, err := this.UserAccessor.Delete(id)
	if err != nil {
		return api.ApiHandlerResponse{Code: http.StatusInternalServerError, Error: err}
	}
	if !res {
		return api.ApiHandlerResponse{Code: http.StatusNotFound, Error: errors.New("data not found")}
	}

	return api.ApiHandlerResponse{Code: http.StatusOK, Body: api.ResponseBody{Message: "data deleted"}}
}

func (this *ApiServiceImpl) AddUser(dto api.AddUserDTO) api.ApiHandlerResponse {
	user, err := this.UserAccessor.Insert(
		dto.Name,
		dto.Email,
	)
	if err != nil {
		return api.ApiHandlerResponse{Code: http.StatusInternalServerError, Error: err}
	}

	return api.ApiHandlerResponse{Code: http.StatusCreated, Body: api.ResponseBody{Data: user}}
}

func (this *ApiServiceImpl) GetProduct(id int) api.ApiHandlerResponse {
	product, err := this.ProductAccessor.Get(id)
	if err != nil {
		return api.ApiHandlerResponse{Code: http.StatusInternalServerError, Error: err}
	}
	if product == nil {
		return api.ApiHandlerResponse{Code: http.StatusNotFound, Error: errors.New("data not found")}
	}

	return api.ApiHandlerResponse{Code: http.StatusOK, Body: api.ResponseBody{Data: product}}
}

func (this *ApiServiceImpl) DeleteProduct(id int) api.ApiHandlerResponse {
	res, err := this.ProductAccessor.Delete(id)
	if err != nil {
		return api.ApiHandlerResponse{Code: http.StatusInternalServerError, Error: err}
	}
	if !res {
		return api.ApiHandlerResponse{Code: http.StatusNotFound, Error: errors.New("data not found")}
	}

	return api.ApiHandlerResponse{Code: http.StatusOK, Body: api.ResponseBody{Message: "data deleted"}}
}

func (this *ApiServiceImpl) AddProduct(dto api.AddProductDTO) api.ApiHandlerResponse {
	product, err := this.ProductAccessor.Insert(
		dto.UserID,
		dto.ProductName,
		dto.ProductDescription,
		dto.ProductPrice,
		dto.ProductImages,
	)
	if err != nil {
		return api.ApiHandlerResponse{Code: http.StatusInternalServerError, Error: err}
	}

	err = this.MQService.Publish(api.MQPublish{
		Queue:   "AddProduct",
		Message: strconv.Itoa(product.Id),
	})
	if err != nil {
		return api.ApiHandlerResponse{Code: http.StatusInternalServerError, Error: err}
	}

	return api.ApiHandlerResponse{Code: http.StatusCreated, Body: api.ResponseBody{Data: product}}
}
