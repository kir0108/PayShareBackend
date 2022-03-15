package main

import (
	"bytes"
	"encoding/json"
	"github.com/kir0108/PayShareBackend/internal/auth_api"
	"github.com/kir0108/PayShareBackend/internal/jwt"
	"net/http"
	"time"
)

const (
	CorrectTestApiId     = "correct_test_api_id"
	IncorrectTestApiId   = "incorrect_test_api_id"
	GoogleApiName        = "google"
	VkApiName            = "vk"
	DefaultCode          = "172637"
	DefaultUserId        = 1
	DefaultRoomId        = 1
	DefaultParticipantId = 1
)

type RequestArgs struct {
	Method      string
	URL         string
	QueryParams map[string]string
	Headers     http.Header
	Body        interface{}
}

type ResponseTest struct {
	Headers http.Header
	Code    int
	Body    string
}

func (r *RequestArgs) GetRequest() (*http.Request, error) {
	body, err := json.Marshal(r.Body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(r.Method, r.URL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	if r.Headers != nil {
		req.Header = r.Headers
	}

	q := req.URL.Query()

	for key, value := range r.QueryParams {
		q.Add(key, value)
	}

	req.URL.RawQuery = q.Encode()

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func GetTestApplication() *application {
	logger, _, _ := newLogger(&config{Production: false})
	manager := jwt.NewManger(&jwt.Config{
		Secret:     "test_secret",
		Expiration: 20 * time.Minute,
	})

	config, err := getConfig()
	if err != nil {
		logger.Fatal(err)
	}

	config.Secret = "test_secret"

	return &application{
		config:        config,
		logger:        logger,
		jwts:          manager,
		codes:         MockCodesRepo{},
		users:         MockUsersRepo{},
		rooms:         MockRoomsRepo{},
		participants:  MockParticipantsRepo{},
		purchases:     MockPurchasesRepo{},
		refreshTokens: MockRefreshTokenRepo{},
		api:           auth_api.NewApi(MockApi{}),
	}
}
