package postgres

import (
	"context"
	"errors"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kir0108/PayShareBackend/internal/data/models"
)

type PurchaseRepo struct {
	DB *pgxpool.Pool
}

func (pr *PurchaseRepo) GetById(ctx context.Context, id int64) (*models.Purchase, error) {
	conn, err := pr.DB.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	defer conn.Release()

	query := "SELECT id, owner_id, room_id, p_name, locate, cost FROM purchases WHERE id=$1"

	purchase := &models.Purchase{}
	if err := pgxscan.Get(ctx, conn, purchase, query, id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrNoRecord
		}

		return nil, err
	}

	return purchase, nil
}

func (pr *PurchaseRepo) Add(ctx context.Context, purchase *models.Purchase) error {
	conn, err := pr.DB.Acquire(ctx)
	if err != nil {
		return err
	}

	defer conn.Release()

	query := "INSERT INTO purchases (owner_id, room_id, p_name, locate, cost) VALUES ($1, $2, $3, $4, $5) " +
		"RETURNING id"
	if err := conn.QueryRow(ctx, query, purchase.OwnerId, purchase.RoomId, purchase.PName, purchase.Locate, purchase.Cost).
		Scan(&purchase.Id); err != nil {

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr); pgErr.Code == pgerrcode.ForeignKeyViolation {
			return models.ErrNoReference
		}

		return err
	}

	return nil
}

func (pr *PurchaseRepo) Update(ctx context.Context, purchase *models.Purchase) error {
	conn, err := pr.DB.Acquire(ctx)
	if err != nil {
		return err
	}

	defer conn.Release()

	query := "UPDATE purchases SET owner_id = $2, room_id = $3, p_name = $4, locate = $5, cost = $6 WHERE id = $1"
	if _, err := conn.Exec(ctx, query, purchase.Id, purchase.OwnerId, purchase.RoomId, purchase.PName, purchase.Locate, purchase.Cost);
	err != nil {

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr); pgErr.Code == pgerrcode.ForeignKeyViolation {
			return models.ErrNoReference
		}

		return err
	}

	return nil
}

func (pr *PurchaseRepo) Delete(ctx context.Context, id int64) error {
	conn, err := pr.DB.Acquire(ctx)
	if err != nil {
		return err
	}

	defer conn.Release()

	query := "DELETE FROM purchases WHERE id = $1"
	if _, err := conn.Exec(ctx, query, id); err != nil {
		return err
	}

	return nil
}
