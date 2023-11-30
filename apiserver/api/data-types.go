package api

type ApiHandlerResponse struct {
	Code  int
	Body  ResponseBody
	Error error
}

type ResponseBody struct {
	Data    any
	Message string
}

type AddUserDTO struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
}

type AddProductDTO struct {
	UserID             int      `json:"user_id" validate:"required"`
	ProductName        string   `json:"product_name" validate:"required"`
	ProductDescription string   `json:"product_description" validate:"required"`
	ProductImages      []string `json:"product_images" validate:"required"`
	ProductPrice       float64  `json:"product_price" validate:"required"`
}

type MQPublish struct {
	Queue   string
	Message string
}

type User struct {
	Id        int
	Name      string
	Email     string
	CreatedAt int64
	UpdatedAt int64
}

type Product struct {
	Id               int
	Name             string
	Description      string
	Images           []string
	Price            float64
	CompressedImages []string `json:"compressed_images"`
	CreatedAt        int64
	UpdatedAt        int64
}
