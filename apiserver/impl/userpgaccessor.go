package impl

import (
	"andrewsaputra/go-message-queue-exercise/apiserver/api"
	"database/sql"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

const COLUMNS_USERS = "id,name,email,created_at,updated_at"

func ConstructUserPGAccessor(addr string) (api.UserDataAccessor, error) {
	cfg, err := pgx.ParseConfig(addr)
	if err != nil {
		return nil, err
	}

	db := stdlib.OpenDB(*cfg)
	return &UserPGAccessor{
		DB: db,
	}, nil
}

type UserPGAccessor struct {
	DB *sql.DB
}

func (this *UserPGAccessor) Get(id int) (*api.User, error) {
	query := fmt.Sprintf(`
		SELECT %s
		FROM users
		WHERE id = $1
	`, COLUMNS_USERS)
	row := this.DB.QueryRow(query, id)
	return this.decode(row)
}

func (this *UserPGAccessor) Insert(name string, email string) (*api.User, error) {
	query := fmt.Sprintf(`
		INSERT INTO users (name,email,created_at)
		VALUES ($1,$2,$3)
		RETURNING %s
	`, COLUMNS_USERS)
	row := this.DB.QueryRow(query, name, email, time.Now().UnixMilli())
	return this.decode(row)
}

func (this *UserPGAccessor) Delete(id int) (bool, error) {
	query := `
		DELETE FROM users
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

func (this *UserPGAccessor) Close() {
	this.DB.Close()
}

func (this *UserPGAccessor) decode(row *sql.Row) (*api.User, error) {
	var data api.User
	if err := row.Scan(&data.Id, &data.Name, &data.Email, &data.CreatedAt, &data.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &data, nil
}
