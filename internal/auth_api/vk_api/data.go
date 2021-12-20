package vk_api

type Response struct {
	Id           int64  `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Photo400Orig string `json:"photo_400_orig"`
	Error        string `json:"error"`
}
