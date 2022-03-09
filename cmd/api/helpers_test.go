package main

import (
	"context"
	"fmt"
	"github.com/kir0108/PayShareBackend/internal/auth_api"
	"github.com/kir0108/PayShareBackend/internal/data/postgres"
	"github.com/kir0108/PayShareBackend/internal/data/redis"
	"github.com/kir0108/PayShareBackend/internal/jwt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
)

func Test_application_generateNumberString(t *testing.T) {
	t.Parallel()
	app := GetTestApplication()

	tests := []struct {
		name      string
		strLen    int
		expStrLen int
	}{
		{
			name:      "generate number string len 0 test",
			strLen:    0,
			expStrLen: 0,
		},
		{
			name:      "generate number string len 5 test",
			strLen:    5,
			expStrLen: 5,
		},
		{
			name:      "generate number string len 7 test",
			strLen:    7,
			expStrLen: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str, err := app.generateNumberString(tt.strLen)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.expStrLen, len(str), str)
			assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("[0-9]{%d}", len(str))), str)
		})
	}
}

func Test_application_generateRandomString(t *testing.T) {
	t.Parallel()
	app := GetTestApplication()
	app.config.Secret = ""

	tests := []struct {
		name      string
		strLen    int
		expStrLen int
	}{
		{
			name:      "generate string len 0 test",
			strLen:    0,
			expStrLen: 0,
		},
		{
			name:      "generate string len 5 test",
			strLen:    5,
			expStrLen: 5,
		},
		{
			name:      "generate string len 7 test",
			strLen:    7,
			expStrLen: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str, err := app.generateRandomString(tt.strLen)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.expStrLen, len(str), str)
			assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("[0-9a-zA-Z]{%d}", len(str))), str)
		})
	}
}

func Test_application_generateRandomUpString(t *testing.T) {
	t.Parallel()
	app := GetTestApplication()

	tests := []struct {
		name      string
		strLen    int
		expStrLen int
	}{
		{
			name:      "generate up string len 0 test",
			strLen:    0,
			expStrLen: 0,
		},
		{
			name:      "generate up string len 5 test",
			strLen:    5,
			expStrLen: 5,
		},
		{
			name:      "generate up string len 7 test",
			strLen:    7,
			expStrLen: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str, err := app.generateRandomUpString(tt.strLen)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.expStrLen, len(str), str)
			assert.Regexp(t, regexp.MustCompile(fmt.Sprintf("[0-9a-zA-Z]{%d}", len(str))), str)
		})
	}
}

func Test_application_getTokens(t *testing.T) {
	t.Parallel()
	app := GetTestApplication()

	tests := []struct {
		name            string
		id              int64
		expAccessToken  string
		expRefreshToken string
	}{
		{
			name:            "generate tokens test",
			id:              1,
			expAccessToken:  "test_access_token",
			expRefreshToken: "test_refresh_token",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokens, err := app.getTokens(context.Background(), tt.id)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.expAccessToken, tokens.AccessToken, tokens.AccessToken)
			assert.Equal(t, tt.expRefreshToken, tokens.RefreshToken, tokens.RefreshToken)
		})
	}
}

func Test_application_writeJSON(t *testing.T) {
	t.Parallel()
	app := GetTestApplication()

	tests := []struct {
		name       string
		status     int
		headers    http.Header
		testStruct interface{}
		expResp    ResponseTest
	}{
		{
			name:       "empty body test",
			status:     200,
			headers:    http.Header{"Content-Type": []string{"application/json"}, "Application": []string{"payshare", "test"}},
			testStruct: nil,
			expResp: ResponseTest{
				Headers: http.Header{"Content-Type": []string{"application/json"}, "Application": []string{"payshare", "test"}},
				Code:    200,
				Body:    "null",
			},
		},
		{
			name:    "empty headers test",
			status:  200,
			headers: http.Header{},
			testStruct: &struct {
				TestData1 int      `json:"test_data_1"`
				TestData2 string   `json:"test_data_2"`
				TestData3 []string `json:"test_data_3"`
				TestData4 *struct {
					TestData1 int    `json:"test_data_1"`
					TestData2 string `json:"test_data_2"`
				} `json:"test_data_4"`
			}{
				TestData1: 25,
				TestData2: "string",
				TestData3: []string{"test1", "test2"},
				TestData4: (*struct {
					TestData1 int    `json:"test_data_1"`
					TestData2 string `json:"test_data_2"`
				})(&struct {
					TestData1 int
					TestData2 string
				}{
					TestData1: 1,
					TestData2: "test2",
				}),
			},
			expResp: ResponseTest{
				Headers: http.Header{"Content-Type": []string{"application/json"}},
				Code:    200,
				Body: `{"test_data_1":25,"test_data_2":"string","test_data_3":["test1","test2"],` +
					`"test_data_4":{"test_data_1":1,"test_data_2":"test2"}}`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			req, err := http.NewRequest("GET", "/test", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				err = app.writeJSON(w, tt.status, tt.testStruct, tt.headers)
			})

			handler.ServeHTTP(rr, req)

			strBody := strings.TrimSpace(rr.Body.String())

			assert.NoError(t, err, err)
			assert.Equal(t, tt.expResp.Code, rr.Code)
			assert.Equal(t, tt.expResp.Body, strBody)
			assert.Equal(t, tt.expResp.Headers, rr.Header())
		})
	}
}

func Test_application_readJSON(t *testing.T) {
	type fields struct {
		config        *config
		logger        *zap.SugaredLogger
		jwts          *jwt.Manager
		codes         redis.CodesRepositoryType
		users         postgres.UserRepoType
		rooms         postgres.RoomRepoType
		participants  postgres.ParticipantRepoType
		purchases     postgres.PurchaseRepoType
		refreshTokens redis.RefreshTokenRepositoryType
		api           *auth_api.Api
	}
	type args struct {
		w   http.ResponseWriter
		r   *http.Request
		dst interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &application{
				config:        tt.fields.config,
				logger:        tt.fields.logger,
				jwts:          tt.fields.jwts,
				codes:         tt.fields.codes,
				users:         tt.fields.users,
				rooms:         tt.fields.rooms,
				participants:  tt.fields.participants,
				purchases:     tt.fields.purchases,
				refreshTokens: tt.fields.refreshTokens,
				api:           tt.fields.api,
			}
			if err := app.readJSON(tt.args.w, tt.args.r, tt.args.dst); (err != nil) != tt.wantErr {
				t.Errorf("readJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
