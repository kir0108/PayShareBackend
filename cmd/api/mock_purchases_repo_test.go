package main

import (
	"context"
	"github.com/kir0108/PayShareBackend/internal/data/models"
)

type MockRoomsRepo struct{}

func (m MockRoomsRepo) Add(ctx context.Context, room *models.Room) error {
	//TODO implement me
	panic("implement me")
}

func (m MockRoomsRepo) UpdateClose(ctx context.Context, roomId int64, close bool) error {
	//TODO implement me
	panic("implement me")
}

func (m MockRoomsRepo) Delete(ctx context.Context, roomId int64) error {
	//TODO implement me
	panic("implement me")
}

func (m MockRoomsRepo) GetById(ctx context.Context, roomId int64) (*models.Room, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockRoomsRepo) GetParticipantOwnerIdById(ctx context.Context, roomId int64) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockRoomsRepo) GetByOwnerId(ctx context.Context, ownerId int64) ([]*models.Room, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockRoomsRepo) GetByUserId(ctx context.Context, userId int64, close bool) ([]*models.Room, error) {
	//TODO implement me
	panic("implement me")
}
