package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/kir0108/PayShareBackend/internal/data/models"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepo struct {
	DB *pgxpool.Pool
}

func (ur *UserRepo) Add(ctx context.Context, user *models.User) error {
	conn, err := ur.DB.Acquire(ctx)
	if err != nil {
		return err
	}

	defer conn.Release()

	query := "INSERT INTO users (api_id, api_name, first_name, second_name, image_url) VALUES ($1, $2, $3, $4, $5) " +
		"RETURNING id"

	if err := conn.QueryRow(ctx, query, user.APIId, user.APIName, user.FirstName, user.SecondName, user.ImageURL).Scan(&user.Id); err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr); pgErr.Code == pgerrcode.UniqueViolation {
			return models.ErrAlreadyExists
		}

		return err
	}

	return nil
}

func (ur *UserRepo) Update(ctx context.Context, user *models.User) error {
	conn, err := ur.DB.Acquire(ctx)
	if err != nil {
		return err
	}

	defer conn.Release()

	query := "UPDATE users SET api_id=$2, api_name=$3, first_name=$4, second_name=$5, image_url=$6 WHERE id = $1"

	if _, err := conn.Exec(ctx, query, user.Id, user.APIId, user.APIId, user.FirstName, user.SecondName, user.ImageURL);
	err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr); pgErr.Code == pgerrcode.UniqueViolation {
			return models.ErrAlreadyExists
		}

		return err
	}

	return nil
}

func (ur *UserRepo) Delete(ctx context.Context, id int64) error {
	conn, err := ur.DB.Acquire(ctx)
	if err != nil {
		return err
	}

	defer conn.Release()

	query := "DELETE FROM users WHERE id = $1"

	tag, err := conn.Exec(ctx, query, id)

	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return models.ErrNoRecord
	}

	return nil
}

func (ur *UserRepo) GetById(ctx context.Context, id int64) (*models.User, error) {
	conn, err := ur.DB.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	defer conn.Release()

	query := "SELECT id, api_id, api_name, first_name, second_name, image_url FROM users WHERE id=$1"

	fmt.Println("ID: ", id)

	user := &models.User{}
	if err := pgxscan.Get(ctx, conn, user, query, id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrNoRecord
		}

		return nil, err
	}

	return user, nil
}

func (ur *UserRepo) GetByAPI(ctx context.Context, apiId string, apiName string) (*models.User, error) {
	fmt.Println(apiId, apiName)
	conn, err := ur.DB.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	defer conn.Release()

	query := "SELECT id, api_id, api_name, first_name, second_name, image_url FROM users WHERE api_id = $1 and api_name = $2"

	user := &models.User{}
	if err := pgxscan.Get(ctx, conn, user, query, apiId, apiName); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrNoRecord
		}

		return nil, err
	}

	return user, nil
}
