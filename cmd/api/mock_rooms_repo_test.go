package main

import (
	"context"
	"github.com/kir0108/PayShareBackend/internal/data/models"
)

type MockPurchasesRepo struct{}

func (m MockPurchasesRepo) GetById(ctx context.Context, id int64) (*models.Purchase, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockPurchasesRepo) GetByRoomId(ctx context.Context, roomId int64) ([]*models.Purchase, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockPurchasesRepo) GetParticipantIdListById(ctx context.Context, purchaseId int64) ([]*models.PurchaseParticipant, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockPurchasesRepo) Add(ctx context.Context, purchase *models.Purchase) error {
	//TODO implement me
	panic("implement me")
}

func (m MockPurchasesRepo) Update(ctx context.Context, purchase *models.Purchase) error {
	//TODO implement me
	panic("implement me")
}

func (m MockPurchasesRepo) Delete(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func (m MockPurchasesRepo) AddParticipantToPurchase(ctx context.Context, purchaseId int64, participantId int64) error {
	//TODO implement me
	panic("implement me")
}

func (m MockPurchasesRepo) UpdatePaidParamPurchase(ctx context.Context, purchaseId int64, participantId int64, paid bool) error {
	//TODO implement me
	panic("implement me")
}

func (m MockPurchasesRepo) DeleteParticipantFromPurchase(ctx context.Context, purchaseId int64, participantId int64) error {
	//TODO implement me
	panic("implement me")
}
