package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kir0108/PayShareBackend/internal/data/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepo_Add(t *testing.T) {
	type fields struct {
		DB *pgxpool.Pool
	}
	type args struct {
		ctx  context.Context
		user *models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ur := &UserRepo{
				DB: tt.fields.DB,
			}
			tt.wantErr(t, ur.Add(tt.args.ctx, tt.args.user), fmt.Sprintf("Add(%v, %v)", tt.args.ctx, tt.args.user))
		})
	}
}

func TestUserRepo_Delete(t *testing.T) {
	type fields struct {
		DB *pgxpool.Pool
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ur := &UserRepo{
				DB: tt.fields.DB,
			}
			tt.wantErr(t, ur.Delete(tt.args.ctx, tt.args.id), fmt.Sprintf("Delete(%v, %v)", tt.args.ctx, tt.args.id))
		})
	}
}

func TestUserRepo_GetByAPI(t *testing.T) {
	type fields struct {
		DB *pgxpool.Pool
	}
	type args struct {
		ctx     context.Context
		apiId   string
		apiName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.User
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ur := &UserRepo{
				DB: tt.fields.DB,
			}
			got, err := ur.GetByAPI(tt.args.ctx, tt.args.apiId, tt.args.apiName)
			if !tt.wantErr(t, err, fmt.Sprintf("GetByAPI(%v, %v, %v)", tt.args.ctx, tt.args.apiId, tt.args.apiName)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetByAPI(%v, %v, %v)", tt.args.ctx, tt.args.apiId, tt.args.apiName)
		})
	}
}

func TestUserRepo_GetById(t *testing.T) {
	type fields struct {
		DB *pgxpool.Pool
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.User
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ur := &UserRepo{
				DB: tt.fields.DB,
			}
			got, err := ur.GetById(tt.args.ctx, tt.args.id)
			if !tt.wantErr(t, err, fmt.Sprintf("GetById(%v, %v)", tt.args.ctx, tt.args.id)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetById(%v, %v)", tt.args.ctx, tt.args.id)
		})
	}
}

func TestUserRepo_Update(t *testing.T) {
	type fields struct {
		DB *pgxpool.Pool
	}
	type args struct {
		ctx  context.Context
		user *models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ur := &UserRepo{
				DB: tt.fields.DB,
			}
			tt.wantErr(t, ur.Update(tt.args.ctx, tt.args.user), fmt.Sprintf("Update(%v, %v)", tt.args.ctx, tt.args.user))
		})
	}
}
