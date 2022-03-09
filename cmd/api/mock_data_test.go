package main

import (
	"bytes"
	"encoding/json"
	"github.com/kir0108/PayShareBackend/internal/auth_api"
	"github.com/kir0108/PayShareBackend/internal/jwt"
	"log"
	"net/http"
	"time"
)

const (
	CorrectTestApiId   = "correct_test_api_id"
	IncorrectTestApiId = "incorrect_test_api_id"
	GoogleApiName      = "google"
	VkApiName          = "vk"
	DefaultCode        = "172637"
	DefaultUserId      = 1
)

type RequestArgs struct {
	Method      string
	URL         string
	QueryParams map[string]string
	Body        interface{}
}

type ResponseTest struct {
	Headers http.Header
	Code    int
	Body    string
}

func FatalErr(name string, err error) {
	logger := &log.Logger{}
	logger.Fatalf("test %s err %s", name, err.Error())
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
