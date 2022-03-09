package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kir0108/PayShareBackend/internal/data/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPurchaseRepo_Add(t *testing.T) {
	type fields struct {
		DB *pgxpool.Pool
	}
	type args struct {
		ctx      context.Context
		purchase *models.Purchase
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
			pr := &PurchaseRepo{
				DB: tt.fields.DB,
			}
			tt.wantErr(t, pr.Add(tt.args.ctx, tt.args.purchase), fmt.Sprintf("Add(%v, %v)", tt.args.ctx, tt.args.purchase))
		})
	}
}

func TestPurchaseRepo_AddParticipantToPurchase(t *testing.T) {
	type fields struct {
		DB *pgxpool.Pool
	}
	type args struct {
		ctx           context.Context
		purchaseId    int64
		participantId int64
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
			pr := &PurchaseRepo{
				DB: tt.fields.DB,
			}
			tt.wantErr(t, pr.AddParticipantToPurchase(tt.args.ctx, tt.args.purchaseId, tt.args.participantId), fmt.Sprintf("AddParticipantToPurchase(%v, %v, %v)", tt.args.ctx, tt.args.purchaseId, tt.args.participantId))
		})
	}
}

func TestPurchaseRepo_Delete(t *testing.T) {
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
			pr := &PurchaseRepo{
				DB: tt.fields.DB,
			}
			tt.wantErr(t, pr.Delete(tt.args.ctx, tt.args.id), fmt.Sprintf("Delete(%v, %v)", tt.args.ctx, tt.args.id))
		})
	}
}

func TestPurchaseRepo_DeleteParticipantFromPurchase(t *testing.T) {
	type fields struct {
		DB *pgxpool.Pool
	}
	type args struct {
		ctx           context.Context
		purchaseId    int64
		participantId int64
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
			pr := &PurchaseRepo{
				DB: tt.fields.DB,
			}
			tt.wantErr(t, pr.DeleteParticipantFromPurchase(tt.args.ctx, tt.args.purchaseId, tt.args.participantId), fmt.Sprintf("DeleteParticipantFromPurchase(%v, %v, %v)", tt.args.ctx, tt.args.purchaseId, tt.args.participantId))
		})
	}
}

func TestPurchaseRepo_GetById(t *testing.T) {
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
		want    *models.Purchase
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := &PurchaseRepo{
				DB: tt.fields.DB,
			}
			got, err := pr.GetById(tt.args.ctx, tt.args.id)
			if !tt.wantErr(t, err, fmt.Sprintf("GetById(%v, %v)", tt.args.ctx, tt.args.id)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetById(%v, %v)", tt.args.ctx, tt.args.id)
		})
	}
}

func TestPurchaseRepo_GetByRoomId(t *testing.T) {
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
		want    []*models.Purchase
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := &PurchaseRepo{
				DB: tt.fields.DB,
			}
			got, err := pr.GetByRoomId(tt.args.ctx, tt.args.roomId)
			if !tt.wantErr(t, err, fmt.Sprintf("GetByRoomId(%v, %v)", tt.args.ctx, tt.args.roomId)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetByRoomId(%v, %v)", tt.args.ctx, tt.args.roomId)
		})
	}
}

func TestPurchaseRepo_GetParticipantIdListById(t *testing.T) {
	type fields struct {
		DB *pgxpool.Pool
	}
	type args struct {
		ctx        context.Context
		purchaseId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.PurchaseParticipant
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := &PurchaseRepo{
				DB: tt.fields.DB,
			}
			got, err := pr.GetParticipantIdListById(tt.args.ctx, tt.args.purchaseId)
			if !tt.wantErr(t, err, fmt.Sprintf("GetParticipantIdListById(%v, %v)", tt.args.ctx, tt.args.purchaseId)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetParticipantIdListById(%v, %v)", tt.args.ctx, tt.args.purchaseId)
		})
	}
}

func TestPurchaseRepo_Update(t *testing.T) {
	type fields struct {
		DB *pgxpool.Pool
	}
	type args struct {
		ctx      context.Context
		purchase *models.Purchase
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
			pr := &PurchaseRepo{
				DB: tt.fields.DB,
			}
			tt.wantErr(t, pr.Update(tt.args.ctx, tt.args.purchase), fmt.Sprintf("Update(%v, %v)", tt.args.ctx, tt.args.purchase))
		})
	}
}

func TestPurchaseRepo_UpdatePaidParamPurchase(t *testing.T) {
	type fields struct {
		DB *pgxpool.Pool
	}
	type args struct {
		ctx           context.Context
		purchaseId    int64
		participantId int64
		paid          bool
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
			pr := &PurchaseRepo{
				DB: tt.fields.DB,
			}
			tt.wantErr(t, pr.UpdatePaidParamPurchase(tt.args.ctx, tt.args.purchaseId, tt.args.participantId, tt.args.paid), fmt.Sprintf("UpdatePaidParamPurchase(%v, %v, %v, %v)", tt.args.ctx, tt.args.purchaseId, tt.args.participantId, tt.args.paid))
		})
	}
}
