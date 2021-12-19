package google_api

import "github.com/kir0108/PayShareBackend/internal/data/models"

type GoogleApi struct {

}

func NewGoogleApi(c *Config) *GoogleApi{
	return &GoogleApi{}
}

func (ga *GoogleApi) GetName() string {
	return "google"
}

func (ga *GoogleApi) GetUser(token string) (*models.User, error) {
	return nil, nil
}
