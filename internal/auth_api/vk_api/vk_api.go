package vk_api

import "github.com/kir0108/PayShareBackend/internal/data/models"

type VKApi struct {

}

func NewVKApi(c *Config) *VKApi{
	return &VKApi{}
}

func (va *VKApi) GetName() string {
	return "vk"
}

func (va *VKApi) GetUser(token string) (*models.User, error) {
	return nil, nil
}
