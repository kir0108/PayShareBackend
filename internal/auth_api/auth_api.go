package auth_api

import (
	"strings"

	"github.com/kir0108/PayShareBackend/internal/data/models"
)

type AuthApi interface {
	GetUser(token string) (*models.User, error)
	GetName() string
}

type Api struct {
	alarmSystems     map[string]AuthApi
	alarmSystemsTime map[string]int64
}

func NewApi(apis ...AuthApi) *Api {
	api := &Api{
		alarmSystems: map[string]AuthApi{},
	}

	for _, v := range apis {
		api.alarmSystems[v.GetName()] = v
	}

	return api
}

func (api *Api) GetAPI(name string) (AuthApi, error) {
	authAPI, ok := api.alarmSystems[strings.ToLower(name)]
	if !ok {
		return nil, models.ErrInvalidAPI
	}

	return authAPI, nil
}
