package api

type ApiService interface {
	GetUser(id int) ApiHandlerResponse
	AddUser(dto AddUserDTO) ApiHandlerResponse
	GetProduct(id int) ApiHandlerResponse
	AddProduct(dto AddProductDTO) ApiHandlerResponse
}

type UserDataAccessor interface {
	Get(id int) (*User, error)
	Insert(name string, email string) (*User, error)
	Delete(id int) (bool, error)
}

type ProductDataAccessor interface {
	Get(id int) (*Product, error)
	Insert(name string, description string, price float64, images []string) (*Product, error)
	Update(id int, updatedAttrs []string, data Product) (*Product, error)
	Delete(id int) (bool, error)
}

type MQService interface {
	Publish(MQPublish) error
	Close()
}
