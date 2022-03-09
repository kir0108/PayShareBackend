package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kir0108/PayShareBackend/internal/data/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParticipantRepo_Add(t *testing.T) {
	type fields struct {
		DB *pgxpool.Pool
	}
	type args struct {
		ctx    context.Context
		userId int64
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
			pr := &ParticipantRepo{
				DB: tt.fields.DB,
			}
			tt.wantErr(t, pr.Add(tt.args.ctx, tt.args.userId, tt.args.roomId), fmt.Sprintf("Add(%v, %v, %v)", tt.args.ctx, tt.args.userId, tt.args.roomId))
		})
	}
}

func TestParticipantRepo_DeleteById(t *testing.T) {
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
			pr := &ParticipantRepo{
				DB: tt.fields.DB,
			}
			tt.wantErr(t, pr.DeleteById(tt.args.ctx, tt.args.id), fmt.Sprintf("DeleteById(%v, %v)", tt.args.ctx, tt.args.id))
		})
	}
}

func TestParticipantRepo_DeleteByUserId(t *testing.T) {
	type fields struct {
		DB *pgxpool.Pool
	}
	type args struct {
		ctx    context.Context
		userId int64
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
			pr := &ParticipantRepo{
				DB: tt.fields.DB,
			}
			tt.wantErr(t, pr.DeleteByUserId(tt.args.ctx, tt.args.userId), fmt.Sprintf("DeleteByUserId(%v, %v)", tt.args.ctx, tt.args.userId))
		})
	}
}

func TestParticipantRepo_Exist(t *testing.T) {
	type fields struct {
		DB *pgxpool.Pool
	}
	type args struct {
		ctx    context.Context
		userId int64
		roomId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := &ParticipantRepo{
				DB: tt.fields.DB,
			}
			got, err := pr.Exist(tt.args.ctx, tt.args.userId, tt.args.roomId)
			if !tt.wantErr(t, err, fmt.Sprintf("Exist(%v, %v, %v)", tt.args.ctx, tt.args.userId, tt.args.roomId)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Exist(%v, %v, %v)", tt.args.ctx, tt.args.userId, tt.args.roomId)
		})
	}
}

func TestParticipantRepo_GetParticipantId(t *testing.T) {
	type fields struct {
		DB *pgxpool.Pool
	}
	type args struct {
		ctx    context.Context
		userId int64
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
			pr := &ParticipantRepo{
				DB: tt.fields.DB,
			}
			got, err := pr.GetParticipantId(tt.args.ctx, tt.args.userId, tt.args.roomId)
			if !tt.wantErr(t, err, fmt.Sprintf("GetParticipantId(%v, %v, %v)", tt.args.ctx, tt.args.userId, tt.args.roomId)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetParticipantId(%v, %v, %v)", tt.args.ctx, tt.args.userId, tt.args.roomId)
		})
	}
}

func TestParticipantRepo_GetParticipantsByRoomId(t *testing.T) {
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
		want    []*models.ParticipantUser
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := &ParticipantRepo{
				DB: tt.fields.DB,
			}
			got, err := pr.GetParticipantsByRoomId(tt.args.ctx, tt.args.roomId)
			if !tt.wantErr(t, err, fmt.Sprintf("GetParticipantsByRoomId(%v, %v)", tt.args.ctx, tt.args.roomId)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetParticipantsByRoomId(%v, %v)", tt.args.ctx, tt.args.roomId)
		})
	}
}
