package vk_api

type ApiResponse struct {
	Response []*struct {
		Id           int64  `json:"id"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		Photo400Orig string `json:"photo_400_orig"`
	} `json:"response"`

	Error *struct {
		ErrorCode     int    `json:"error_code"`
		ErrorMsg      string `json:"error_msg"`
	} `json:"error"`
}
