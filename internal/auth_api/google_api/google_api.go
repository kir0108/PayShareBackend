package google_api

import (
	"encoding/json"
	"errors"
	"github.com/kir0108/PayShareBackend/internal/data/models"
	"net/http"
	"time"
)

type GoogleApi struct {
	httpClient *http.Client
}

func NewGoogleApi(c *Config) *GoogleApi{
	return &GoogleApi{
		httpClient: &http.Client{
			Timeout: time.Second * 2,
		},
	}
}

func (ga *GoogleApi) GetName() string {
	return "google"
}

func (ga *GoogleApi) GetUser(token string) (*models.User, error) {
	req, err := http.NewRequest("GET", "https://oauth2.googleapis.com/tokeninfo", nil)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("id_token", token)
	req.URL.RawQuery = q.Encode()

	resp, err := ga.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var respData Response

	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}

	if respData.Error != "" {
		return nil, errors.New(respData.Error)
	}

	return &models.User{
		APIId:      respData.Sub,
		APIName:    ga.GetName(),
		FirstName:  respData.GivenName,
		SecondName: respData.FamilyName,
		ImageURL:   respData.Picture,
	}, nil
}
