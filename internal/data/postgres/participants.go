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

type ParticipantRepo struct {
	DB *pgxpool.Pool
}

func (pr *ParticipantRepo) Add(ctx context.Context, userId int64, roomId int64) error {
	conn, err := pr.DB.Acquire(ctx)
	if err != nil {
		return err
	}

	defer conn.Release()

	query := "INSERT INTO participants (user_id, room_id) VALUES ($1, $2)"
	if _, err := conn.Exec(ctx, query, userId, roomId); err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr); pgErr.Code == pgerrcode.UniqueViolation {
			return models.ErrAlreadyExists
		}

		return err
	}

	return nil
}

func (pr *ParticipantRepo) Delete(ctx context.Context, userId int64) error {
	conn, err := pr.DB.Acquire(ctx)
	if err != nil {
		return err
	}

	defer conn.Release()

	query := "DELETE FROM participants WHERE user_id = $1"
	if _, err := conn.Exec(ctx, query, userId); err != nil {
		return err
	}

	return nil
}

func (pr *ParticipantRepo) Exist(ctx context.Context, userId int64, roomId int64) (bool, error) {
	conn, err := pr.DB.Acquire(ctx)
	if err != nil {
		return false, err
	}

	defer conn.Release()

	participant := &models.Participant{}

	query := "SELECT user_id, room_id from participants WHERE user_id = $1 and room_id = $2"

	if err := pgxscan.Get(ctx, conn, participant, query, userId, roomId); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
