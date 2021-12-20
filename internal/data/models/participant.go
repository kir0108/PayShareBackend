package models

type Participant struct {
	Id     int64 `json:"id" db:"id"`
	UserId int64 `json:"-" db:"user_id"`
	RoomId int64 `json:"-" db:"room_id"`
}
