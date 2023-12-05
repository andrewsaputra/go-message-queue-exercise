package api

type ApiService interface {
	GetUser(id int) ApiHandlerResponse
	DeleteUser(id int) ApiHandlerResponse
	AddUser(dto AddUserDTO) ApiHandlerResponse
	GetProduct(id int) ApiHandlerResponse
	DeleteProduct(id int) ApiHandlerResponse
	AddProduct(dto AddProductDTO) ApiHandlerResponse
}

type UserDataAccessor interface {
	Get(id int) (*User, error)
	Insert(name string, email string) (*User, error)
	Delete(id int) (bool, error)
	Close()
}

type ProductDataAccessor interface {
	Get(id int) (*Product, error)
	Insert(userId int, name string, description string, price float64, images []string) (*Product, error)
	Delete(id int) (bool, error)
	Close()
}

type MQService interface {
	Publish(MQPublish) error
	Close()
}
