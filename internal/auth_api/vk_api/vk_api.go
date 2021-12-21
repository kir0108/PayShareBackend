package vk_api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/kir0108/PayShareBackend/internal/data/models"
)

type VKApi struct {
	httpClient *http.Client
}

func NewVKApi(c *Config) *VKApi {
	return &VKApi{
		httpClient: &http.Client{
			Timeout: time.Second * 2,
		},
	}
}

func (va *VKApi) GetName() string {
	return "vk"
}

func (va *VKApi) GetUser(token string) (*models.User, error) {
	req, err := http.NewRequest("GET", "https://api.vk.com/method/users.get", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("access_token", token)
	q.Add("fields", "photo_400_orig")
	q.Add("v", "5.131")
	req.URL.RawQuery = q.Encode()

	resp, err := va.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var respData ApiResponse

	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}

	if len(respData.Response) < 1 {
		return nil, errors.New("nil response")
	}

	if respData.Error != nil {
		return nil, errors.New(respData.Error.ErrorMsg)
	}

	return &models.User{
		APIId:      strconv.FormatInt(respData.Response[0].Id, 10),
		APIName:    va.GetName(),
		FirstName:  respData.Response[0].FirstName,
		SecondName: respData.Response[0].LastName,
		ImageURL:   respData.Response[0].Photo400Orig,
	}, nil
}
