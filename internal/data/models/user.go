package models

const (
	GoogleAPIName = "google"
	VKAPIName     = "vk"
)

type User struct {
	Id         int64  `json:"-" db:"id"`
	APIId      string `json:"-" db:"id"`
	APIName    string `json:"-" db:"api_name"`
	FirstName  string `json:"first_name" db:"first_name"`
	SecondName string `json:"second_name" db:"second_name"`
	ImageURL   string `json:"image_url" db:"image_url"`
}
