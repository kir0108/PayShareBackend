package google_api

type Response struct {
	Sub        string `json:"sub"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Picture    string `json:"picture"`
	Error      string `json:"error"`
}
