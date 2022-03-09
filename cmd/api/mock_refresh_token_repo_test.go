package main

import (
	"context"
	"errors"
	"github.com/kir0108/PayShareBackend/internal/data/models"
)

type MockRefreshTokenRepo struct{}

const invalidTokenNotExist = "invalid_refresh_token_not_exist"
const invalidTokenOtherErr = "invalid_token_other_err"

func (m MockRefreshTokenRepo) Add(ctx context.Context, session string, id int64) error {
	switch {
	default:
		return nil
	}
}

func (m MockRefreshTokenRepo) Get(ctx context.Context, session string) (int64, error) {
	switch {
	default:
		return DefaultUserId, nil
	}
}

func (m MockRefreshTokenRepo) Refresh(ctx context.Context, old, new string) error {
	switch old {
	case invalidTokenOtherErr:
		return errors.New("other error")
	case invalidTokenNotExist:
		return models.ErrNoRecord
	default:
		return nil
	}
}

func (m MockRefreshTokenRepo) Delete(ctx context.Context, session string) error {
	switch session {
	case invalidTokenOtherErr:
		return errors.New("other error")
	case invalidTokenNotExist:
		return models.ErrNoRecord
	default:
		return nil
	}
}

func (m MockRefreshTokenRepo) DeleteExpired(ctx context.Context) error {
	switch {
	default:
		return nil
	}
}

func (m MockRefreshTokenRepo) DeleteByUserID(ctx context.Context, id int64) error {
	switch {
	default:
		return nil
	}
}
