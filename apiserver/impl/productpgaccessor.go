package impl

import (
	"andrewsaputra/go-message-queue-exercise/apiserver/api"
	"database/sql"
	"fmt"
	"time"

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

func (this *ProductPGAccessor) Insert(userId int, name string, description string, price float64, images []string) (*api.Product, error) {
	query := fmt.Sprintf(`
		INSERT INTO products (user_id,name,description,images,price,created_at)
		VALUES ($1,$2,$3,$4,$5,$6)
		RETURNING %s
	`, COLUMNS_PRODUCTS)
	row := this.DB.QueryRow(query, userId, name, description, pq.Array(images), price, time.Now().UnixMilli())
	return this.decode(row)
}

func (this *ProductPGAccessor) Delete(id int) (bool, error) {
	query := `
		DELETE FROM products
		WHERE id = $1
	`
	result, err := this.DB.Exec(query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	numRows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return numRows == 1, nil
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
