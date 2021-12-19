package postgres

import (
	"context"
	"errors"
	"github.com/kir0108/PayShareBackend/internal/data/models"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type RoomRepo struct {
	DB *pgxpool.Pool
}

func (rr *RoomRepo) Add(ctx context.Context, room *models.Room) error {
	conn, err := rr.DB.Acquire(ctx)
	if err != nil {
		return err
	}

	defer conn.Release()

	query := "INSERT INTO rooms (owner_id, room_name, room_date) VALUES ($1, $2, $3) " +
		"RETURNING id"

	if err := conn.QueryRow(ctx, query, room.OwnerId, room.RoomName, room.RoomDate).Scan(&room.Id); err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr); pgErr.Code == pgerrcode.UniqueViolation {
			return models.ErrAlreadyExists
		}

		return err
	}

	return nil
}

func (rr *RoomRepo) UpdateClose(ctx context.Context, roomId int64, close bool) error {
	conn, err := rr.DB.Acquire(ctx)
	if err != nil {
		return err
	}

	defer conn.Release()

	query := "UPDATE rooms SET close=$2 WHERE id = $1"

	_, err = conn.Exec(ctx, query, roomId, close)

	if err != nil {
		return err
	}

	return nil
}

func (rr *RoomRepo) Delete(ctx context.Context, roomId int64) error {
	conn, err := rr.DB.Acquire(ctx)
	if err != nil {
		return err
	}

	defer conn.Release()

	query := "DELETE FROM rooms WHERE id = $1"

	tag, err := conn.Exec(ctx, query, roomId)

	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return models.ErrNoRecord
	}

	return nil
}

func (rr *RoomRepo) GetById(ctx context.Context, roomId int64) (*models.Room, error) {
	conn, err := rr.DB.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	defer conn.Release()

	query := "SELECT id, owner_id, room_name, room_date, close FROM rooms WHERE id=$1"

	room := &models.Room{}
	if err := pgxscan.Get(ctx, conn, room, query, roomId); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrNoRecord
		}

		return nil, err
	}

	return room, nil
}

func (rr *RoomRepo) GetByOwnerId(ctx context.Context, ownerId int64) ([]*models.Room, error) {
	conn, err := rr.DB.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	defer conn.Release()

	query := "SELECT id, owner_id, room_name, room_date, close FROM rooms WHERE owner_id=$1"

	var rooms []*models.Room

	if err := pgxscan.Select(ctx, conn, &rooms, query, ownerId); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrNoRecord
		}

		return nil, err
	}

	return rooms, nil
}
