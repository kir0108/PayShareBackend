package models

type Room struct {
	Id       int64  `json:"id" db:"id"`
	OwnerId  int64  `json:"-" db:"owner_id"`
	RoomName string `json:"room_name" db:"room_name"`
	RoomDate string `json:"room_date" db:"room_date"`
	Close    bool   `json:"close" json:"close"`
}

type RoomElement struct {
	Room      *Room       `json:"room"`
	Purchases []*Purchase `json:"purchases"`
	IsYour    bool        `json:"is_your"`
}

type RoomInfo struct {
	OwnerParticipantId int64              `json:"owner_participant_id"`
	Participants       []*ParticipantUser `json:"participants"`
	Purchases          []*PurchaseRoom    `json:"purchases"`
}
