package postgres

import (
	"context"
	"database/sql"
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

func (pr *ParticipantRepo) GetParticipantId(ctx context.Context, userId int64, roomId int64) (int64, error) {
	conn, err := pr.DB.Acquire(ctx)
	if err != nil {
		return 0, err
	}

	defer conn.Release()

	var id sql.NullInt64

	query := "SELECT id FROM participants WHERE user_id = $1 and room_id = $2"
	if err := conn.QueryRow(ctx, query, userId, roomId).Scan(&id); err != nil {
		return 0, err
	}

	if !id.Valid {
		return 0, models.ErrNoRecord
	}

	return id.Int64, nil
}

func (pr *ParticipantRepo) GetParticipantsByRoomId(ctx context.Context, roomId int64) ([]*models.ParticipantUser, error) {
	conn, err := pr.DB.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	defer conn.Release()

	var participants []*models.ParticipantUser

	query := "select p.id as id, first_name, second_name, image_url from participants as p " +
		"join users u on p.user_id = u.id where p.room_id = $1;"
	if err := pgxscan.Select(ctx, conn, &participants, query, roomId); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrNoRecord
		}

		return nil, err
	}

	return participants, nil
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

func (pr *ParticipantRepo) DeleteById(ctx context.Context, id int64) error {
	conn, err := pr.DB.Acquire(ctx)
	if err != nil {
		return err
	}

	defer conn.Release()

	query := "DELETE FROM participants WHERE id = $1"
	if _, err := conn.Exec(ctx, query, id); err != nil {
		return err
	}

	return nil
}

func (pr *ParticipantRepo) DeleteByUserId(ctx context.Context, userId int64) error {
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
