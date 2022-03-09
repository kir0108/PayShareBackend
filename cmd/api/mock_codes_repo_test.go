package main

import "context"

type MockCodesRepo struct{}

func (m MockCodesRepo) Add(ctx context.Context, code string, id int64) error {
	switch {
	default:
		return nil
	}
}

func (m MockCodesRepo) GetCode(ctx context.Context, id int64) (string, error) {
	switch {
	default:
		return DefaultCode, nil
	}
}

func (m MockCodesRepo) GetId(ctx context.Context, code string) (int64, error) {
	switch {
	default:
		return DefaultUserId, nil
	}
}

func (m MockCodesRepo) Exist(ctx context.Context, code string) (bool, error) {
	switch {
	default:
		return true, nil
	}
}

func (m MockCodesRepo) CodeExpired(ctx context.Context, code string) (bool, error) {
	switch {
	default:
		return true, nil
	}
}

func (m MockCodesRepo) Delete(ctx context.Context, code string) error {
	switch {
	default:
		return nil
	}
}
