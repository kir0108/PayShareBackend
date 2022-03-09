package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kir0108/PayShareBackend/internal/data/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRoomRepo_Add(t *testing.T) {
	type fields struct {
		DB *pgxpool.Pool
	}
	type args struct {
		ctx  context.Context
		room *models.Room
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
			rr := &RoomRepo{
				DB: tt.fields.DB,
			}
			tt.wantErr(t, rr.Add(tt.args.ctx, tt.args.room), fmt.Sprintf("Add(%v, %v)", tt.args.ctx, tt.args.room))
		})
	}
}

func TestRoomRepo_Delete(t *testing.T) {
	type fields struct {
		DB *pgxpool.Pool
	}
	type args struct {
		ctx    context.Context
		roomId int64
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
			rr := &RoomRepo{
				DB: tt.fields.DB,
			}
			tt.wantErr(t, rr.Delete(tt.args.ctx, tt.args.roomId), fmt.Sprintf("Delete(%v, %v)", tt.args.ctx, tt.args.roomId))
		})
	}
}

func TestRoomRepo_GetById(t *testing.T) {
	type fields struct {
		DB *pgxpool.Pool
	}
	type args struct {
		ctx    context.Context
		roomId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Room
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := &RoomRepo{
				DB: tt.fields.DB,
			}
			got, err := rr.GetById(tt.args.ctx, tt.args.roomId)
			if !tt.wantErr(t, err, fmt.Sprintf("GetById(%v, %v)", tt.args.ctx, tt.args.roomId)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetById(%v, %v)", tt.args.ctx, tt.args.roomId)
		})
	}
}

func TestRoomRepo_GetByOwnerId(t *testing.T) {
	type fields struct {
		DB *pgxpool.Pool
	}
	type args struct {
		ctx     context.Context
		ownerId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.Room
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := &RoomRepo{
				DB: tt.fields.DB,
			}
			got, err := rr.GetByOwnerId(tt.args.ctx, tt.args.ownerId)
			if !tt.wantErr(t, err, fmt.Sprintf("GetByOwnerId(%v, %v)", tt.args.ctx, tt.args.ownerId)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetByOwnerId(%v, %v)", tt.args.ctx, tt.args.ownerId)
		})
	}
}

func TestRoomRepo_GetByUserId(t *testing.T) {
	type fields struct {
		DB *pgxpool.Pool
	}
	type args struct {
		ctx    context.Context
		userId int64
		close  bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.Room
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := &RoomRepo{
				DB: tt.fields.DB,
			}
			got, err := rr.GetByUserId(tt.args.ctx, tt.args.userId, tt.args.close)
			if !tt.wantErr(t, err, fmt.Sprintf("GetByUserId(%v, %v, %v)", tt.args.ctx, tt.args.userId, tt.args.close)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetByUserId(%v, %v, %v)", tt.args.ctx, tt.args.userId, tt.args.close)
		})
	}
}

func TestRoomRepo_GetParticipantOwnerIdById(t *testing.T) {
	type fields struct {
		DB *pgxpool.Pool
	}
	type args struct {
		ctx    context.Context
		roomId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := &RoomRepo{
				DB: tt.fields.DB,
			}
			got, err := rr.GetParticipantOwnerIdById(tt.args.ctx, tt.args.roomId)
			if !tt.wantErr(t, err, fmt.Sprintf("GetParticipantOwnerIdById(%v, %v)", tt.args.ctx, tt.args.roomId)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetParticipantOwnerIdById(%v, %v)", tt.args.ctx, tt.args.roomId)
		})
	}
}

func TestRoomRepo_UpdateClose(t *testing.T) {
	type fields struct {
		DB *pgxpool.Pool
	}
	type args struct {
		ctx    context.Context
		roomId int64
		close  bool
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
			rr := &RoomRepo{
				DB: tt.fields.DB,
			}
			tt.wantErr(t, rr.UpdateClose(tt.args.ctx, tt.args.roomId, tt.args.close), fmt.Sprintf("UpdateClose(%v, %v, %v)", tt.args.ctx, tt.args.roomId, tt.args.close))
		})
	}
}
