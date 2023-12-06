package api

type MQService interface {
	Subscribe(queueName string, consumer ConsumerService) error
	Close()
}

type ConsumerService interface {
	OnConsumed(body []byte) error
}

type ProductDataAccessor interface {
	Get(id int) (*Product, error)
	SetCompressedImages(id int, urls []string) (*Product, error)
	Close()
}
