package main

import (
	"context"
	"github.com/kir0108/PayShareBackend/internal/data/models"
)

type MockParticipantsRepo struct{}

func (m MockParticipantsRepo) GetParticipantId(ctx context.Context, userId int64, roomId int64) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockParticipantsRepo) GetParticipantsByRoomId(ctx context.Context, roomId int64) ([]*models.ParticipantUser, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockParticipantsRepo) Add(ctx context.Context, userId int64, roomId int64) error {
	//TODO implement me
	panic("implement me")
}

func (m MockParticipantsRepo) DeleteById(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func (m MockParticipantsRepo) DeleteByUserId(ctx context.Context, userId int64) error {
	//TODO implement me
	panic("implement me")
}

func (m MockParticipantsRepo) Exist(ctx context.Context, userId int64, roomId int64) (bool, error) {
	//TODO implement me
	panic("implement me")
}
