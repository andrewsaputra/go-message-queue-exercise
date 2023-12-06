package api

type Product struct {
	Id               int
	UserId           int
	Name             string
	Description      string
	Images           []string
	Price            float64
	CompressedImages []string
	CreatedAt        int64
	UpdatedAt        int64
}
