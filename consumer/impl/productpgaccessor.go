package impl

import (
	"andrewsaputra/go-message-queue-exercise/consumer/api"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/lib/pq"
)

const COLUMNS_PRODUCTS = "id,user_id,name,description,images,compressed_images,price,created_at,updated_at"

func ConstructProductPGAccessor(addr string) (api.ProductDataAccessor, error) {
	cfg, err := pgx.ParseConfig(addr)
	if err != nil {
		return nil, err
	}

	db := stdlib.OpenDB(*cfg)
	return &ProductPGAccessor{
		DB: db,
	}, nil
}

type ProductPGAccessor struct {
	DB *sql.DB
}

func (this *ProductPGAccessor) Get(id int) (*api.Product, error) {
	query := fmt.Sprintf(`
		SELECT %s
		FROM products
		WHERE id = $1
	`, COLUMNS_PRODUCTS)
	row := this.DB.QueryRow(query, id)
	return this.decode(row)
}

func (this *ProductPGAccessor) SetCompressedImages(id int, urls []string) (*api.Product, error) {
	query := fmt.Sprintf(`
		UPDATE products
		SET compressed_images = $1
		WHERE id = $2
		RETURNING %s
	`, COLUMNS_PRODUCTS)
	row := this.DB.QueryRow(query, pq.Array(urls), id)
	return this.decode(row)
}

func (this *ProductPGAccessor) Close() {
	this.DB.Close()
}

func (this *ProductPGAccessor) decode(row *sql.Row) (*api.Product, error) {
	var product api.Product
	if err := row.Scan(
		&product.Id, &product.UserId, &product.Name, &product.Description,
		pq.Array(&product.Images), pq.Array(&product.CompressedImages), &product.Price,
		&product.CreatedAt, &product.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &product, nil
}
