package main

import (
	"context"
	"errors"
	"github.com/kir0108/PayShareBackend/internal/data/models"
)

const (
	ErrOtherRoomByName = "Jdhqodhal"
	ErrOtherRoomById   = 1847
	ErrNoExistRoomById = 1988
)

type MockRoomsRepo struct{}

func (m MockRoomsRepo) Add(ctx context.Context, room *models.Room) error {
	switch {
	case room.RoomName == ErrOtherRoomByName:
		return errors.New("test")
	default:
		room.Id = DefaultRoomId
		room.OwnerId = DefaultUserId
		room.Close = false
		return nil
	}
}

func (m MockRoomsRepo) UpdateClose(ctx context.Context, roomId int64, close bool) error {
	//TODO implement me
	panic("implement me")
}

func (m MockRoomsRepo) Delete(ctx context.Context, roomId int64) error {
	switch roomId {
	case ErrOtherRoomById:
		return errors.New("test")
	case ErrNoExistRoomById:
		return models.ErrNoRecord
	default:
		return nil
	}
}

func (m MockRoomsRepo) GetById(ctx context.Context, roomId int64) (*models.Room, error) {
	switch roomId {
	case ErrOtherRoomById:
		return nil, errors.New("test")
	case ErrNoExistRoomById:
		return nil, models.ErrNoRecord
	default:
		return &models.Room{
			Id:       DefaultRoomId,
			OwnerId:  DefaultUserId,
			RoomName: "Test",
			RoomDate: "Date",
			Close:    false,
		}, nil
	}
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
	switch {
	case (userId == DefaultUserId) && close:
		return []*models.Room{
			{
				Id:       1,
				OwnerId:  DefaultUserId,
				RoomName: "test",
				RoomDate: "test",
				Close:    true,
			},
			{
				Id:       2,
				OwnerId:  3,
				RoomName: "test",
				RoomDate: "test",
				Close:    true,
			},
		}, nil
	case (userId == DefaultUserId) && !close:
		return []*models.Room{
			{
				Id:       1,
				OwnerId:  DefaultUserId,
				RoomName: "test",
				RoomDate: "test",
				Close:    false,
			},
			{
				Id:       2,
				OwnerId:  3,
				RoomName: "test",
				RoomDate: "test",
				Close:    false,
			},
		}, nil
	default:
		return nil, nil
	}
}
