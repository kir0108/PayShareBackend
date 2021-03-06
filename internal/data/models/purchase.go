package models

type Locate struct {
	Lat         float64 `json:"lat" db:"lat"`
	Long        float64 `json:"long" db:"long"`
	ShopName    string  `json:"shop_name" db:"shop_name"`
	Date        string  `json:"date" db:"date"`
	Description string  `json:"description" db:"description"`
}

type Purchase struct {
	Id      int64   `json:"id" db:"id"`
	OwnerId int64   `json:"owner_id" db:"owner_id"`
	RoomId  int64   `json:"-" db:"room_id"`
	PName   string  `json:"name" db:"p_name"`
	Locate  *Locate `json:"locate" db:"locate"`
	Cost    int64   `json:"cost" db:"cost"`
}

type PurchaseParticipant struct {
	ParticipantId int64 `json:"participant_id" db:"participant_id"`
	Paid          bool  `json:"paid" db:"paid"`
}

type PurchaseRoom struct {
	*Purchase
	Participants []*PurchaseParticipant `json:"participants"`
}
