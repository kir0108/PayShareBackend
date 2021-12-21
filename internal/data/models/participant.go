package models

type Participant struct {
	Id     int64 `json:"id" db:"id"`
	UserId int64 `json:"-" db:"user_id"`
	RoomId int64 `json:"-" db:"room_id"`
}

type ParticipantUser struct {
	Id         int64  `json:"id"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	ImageURL   string `json:"image_url"`
}
