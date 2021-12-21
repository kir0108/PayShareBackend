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

func (pr *PurchaseRepo) GetByRoomId(ctx context.Context, roomId int64) ([]*models.Purchase, error) {
	conn, err := pr.DB.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	defer conn.Release()

	query := "SELECT id, owner_id, room_id, p_name, locate, cost FROM purchases WHERE room_id=$1"

	var purchases []*models.Purchase

	if err := pgxscan.Select(ctx, conn, &purchases, query, roomId); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrNoRecord
		}

		return nil, err
	}

	return purchases, nil
}

func (pr *PurchaseRepo) GetParticipantIdListById(ctx context.Context, purchaseId int64) ([]*models.PurchaseParticipant, error) {
	conn, err := pr.DB.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	defer conn.Release()

	query := "SELECT participant_id, paid FROM participants_purchases WHERE purchase_id=$1"

	var participants []*models.PurchaseParticipant

	if err := pgxscan.Select(ctx, conn, &participants, query, purchaseId); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrNoRecord
		}

		return nil, err
	}

	return participants, nil
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

func (pr *PurchaseRepo) AddParticipantToPurchase(ctx context.Context, purchaseId int64, participantId int64) error {
	conn, err := pr.DB.Acquire(ctx)
	if err != nil {
		return err
	}

	defer conn.Release()

	query := "INSERT INTO participants_purchases (purchase_id, participant_id) VALUES ($1, $2)"
	if _, err := conn.Exec(ctx, query, purchaseId, participantId);
		err != nil {

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr); pgErr.Code == pgerrcode.ForeignKeyViolation {
			return models.ErrNoReference
		}

		if errors.As(err, &pgErr); pgErr.Code == pgerrcode.UniqueViolation {
			return models.ErrAlreadyExists
		}

		return err
	}

	return nil
}

func (pr *PurchaseRepo) UpdatePaidParamPurchase(ctx context.Context, purchaseId int64, participantId int64, paid bool) error {
	conn, err := pr.DB.Acquire(ctx)
	if err != nil {
		return err
	}

	defer conn.Release()

	query := "UPDATE participants_purchases SET paid = $3 WHERE purchase_id = $1 and participant_id = $2"
	if _, err := conn.Exec(ctx, query, purchaseId, participantId, paid); err != nil {
		return err
	}

	return nil
}

func (pr *PurchaseRepo) DeleteParticipantFromPurchase(ctx context.Context, purchaseId int64, participantId int64) error {
	conn, err := pr.DB.Acquire(ctx)
	if err != nil {
		return err
	}

	defer conn.Release()

	query := "DELETE FROM participants_purchases WHERE purchase_id = $1 and participant_id = $2"
	if _, err := conn.Exec(ctx, query, purchaseId, participantId);
		err != nil {
		return err
	}

	return nil
}
