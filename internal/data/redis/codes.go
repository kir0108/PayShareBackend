package redis

import (
	"context"
	"errors"
	"github.com/kir0108/PayShareBackend/internal/data/models"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

type CodesRepository struct {
	Pool           *redis.Pool
	CodeHideTTL    time.Duration
	CodeExpiredTTL time.Duration
}

func NewCodesRepository(pool *redis.Pool, c *Config) *CodesRepository {
	return &CodesRepository{
		Pool:           pool,
		CodeHideTTL:    c.CodeHideTTL,
		CodeExpiredTTL: c.CodeExpiredTTL,
	}
}

func (cr *CodesRepository) Add(ctx context.Context, code string, id int64) error {
	conn, err := cr.Pool.GetContext(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	if err := conn.Send("SET", "code:"+code, id, "EX", int(cr.CodeExpiredTTL.Seconds())); err != nil {
		return err
	}

	if err := conn.Send("SET", "id_code:"+strconv.FormatInt(id, 10), code, "EX", int(cr.CodeHideTTL.Seconds())); err != nil {
		if err := conn.Send("DEL", "code:"+code); err != nil {
			return err
		}

		return err
	}

	return nil
}

func (cr *CodesRepository) GetCode(ctx context.Context, id int64) (string, error) {
	conn, err := cr.Pool.GetContext(ctx)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	code, err := redis.String(conn.Do("GET", "id_code:"+strconv.FormatInt(id, 10)))
	if err != nil {
		if errors.Is(err, redis.ErrNil) {
			return "", models.ErrNoRecord
		}

		return "", err
	}

	return code, nil
}

func (cr *CodesRepository) GetId(ctx context.Context, code string) (int64, error) {
	conn, err := cr.Pool.GetContext(ctx)
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	id, err := redis.Int64(conn.Do("GET", "code:"+code))
	if err != nil {
		if errors.Is(err, redis.ErrNil) {
			return 0, models.ErrNoRecord
		}
		return 0, err
	}

	return id, nil
}

func (cr *CodesRepository) Exist(ctx context.Context, code string) (bool, error) {
	conn, err := cr.Pool.GetContext(ctx)
	if err != nil {
		return false, err
	}
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", "code:"+code))
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (cr *CodesRepository) CodeExpired(ctx context.Context, code string) (bool, error) {
	conn, err := cr.Pool.GetContext(ctx)
	if err != nil {
		return false, err
	}
	defer conn.Close()

	ttl, err := redis.Int64(conn.Do("TTL", "code:"+code))
	if err != nil {
		return false, err
	}

	if ttl < int64(cr.CodeExpiredTTL - cr.CodeHideTTL) {
		return true, nil
	}

	return false, nil
}

func (cr *CodesRepository) Delete(ctx context.Context, code string) error {
	conn, err := cr.Pool.GetContext(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	return conn.Send("DEL", "code:"+code)
}
