package main

import (
	"fmt"
	"github.com/kir0108/PayShareBackend/internal/data/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_application_loginHandler(t *testing.T) {
	t.Parallel()
	app := GetTestApplication()
	handler := http.HandlerFunc(app.loginHandler)

	type loginBody struct {
		Api   string `json:"auth_api"`
		Token string `json:"token"`
	}

	tests := []struct {
		name    string
		reqArgs RequestArgs
		expResp ResponseTest
	}{
		{
			name: "correct login test",
			reqArgs: RequestArgs{
				Method:      "POST",
				URL:         "/auth/login",
				QueryParams: nil,
				Body: &loginBody{
					Api: "test",
					Token: TokenCreate(models.User{
						APIId:      "135234",
						APIName:    "test",
						FirstName:  "Daniil",
						SecondName: "Firsov",
					}),
				},
			},
			expResp: ResponseTest{
				Code: 200,
				Body: `{"access_token":"test_access_token","refresh_token":"test_refresh_token"}`,
			},
		},
		{
			name: "invalid api name test",
			reqArgs: RequestArgs{
				Method:      "POST",
				URL:         "/auth/login",
				QueryParams: nil,
				Body: &loginBody{
					Api: "test123",
					Token: TokenCreate(models.User{
						APIId:      "135234",
						APIName:    "test123",
						FirstName:  "Daniil",
						SecondName: "Firsov",
					}),
				},
			},
			expResp: ResponseTest{
				Code: 400,
				Body: `{"error":"invalid api name"}`,
			},
		},
		{
			name: "invalid api token test",
			reqArgs: RequestArgs{
				Method:      "POST",
				URL:         "/auth/login",
				QueryParams: nil,
				Body: &loginBody{
					Api:   "test",
					Token: "invalid_token",
				},
			},
			expResp: ResponseTest{
				Code: 500,
				Body: `{"error":"the server encountered a problem and could not process your request"}`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := tt.reqArgs.GetRequest()
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			strings.TrimSpace(rr.Body.String())

			assert.Equal(t, tt.expResp, ResponseTest{
				Code: rr.Code,
				Body: strings.TrimSpace(rr.Body.String()),
			})
		})
	}
}

func Test_application_logoutUserHandler(t *testing.T) {
	t.Parallel()
	app := GetTestApplication()
	handler := http.HandlerFunc(app.logoutUserHandler)

	type requestBody struct {
		RefreshToken string `json:"refresh_token"`
	}

	tests := []struct {
		name    string
		reqArgs RequestArgs
		expResp ResponseTest
	}{
		{
			name: "correct login test",
			reqArgs: RequestArgs{
				Method:      "POST",
				URL:         "/auth/login",
				QueryParams: nil,
				Body: &requestBody{
					RefreshToken: "qwert",
				},
			},
			expResp: ResponseTest{
				Code: 200,
				Body: ``,
			},
		},
		{
			name: "token not exist test",
			reqArgs: RequestArgs{
				Method:      "POST",
				URL:         "/auth/login",
				QueryParams: nil,
				Body:        &requestBody{RefreshToken: invalidTokenNotExist},
			},
			expResp: ResponseTest{
				Code: 401,
				Body: fmt.Sprintf(`{"error":"no such session"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := tt.reqArgs.GetRequest()
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			strings.TrimSpace(rr.Body.String())

			assert.Equal(t, tt.expResp, ResponseTest{
				Code: rr.Code,
				Body: strings.TrimSpace(rr.Body.String()),
			})
		})
	}
}

func Test_application_refreshTokenHandler(t *testing.T) {
	t.Parallel()
	app := GetTestApplication()
	handler := http.HandlerFunc(app.refreshTokenHandler)

	type requestBody struct {
		RefreshToken string `json:"refresh_token"`
	}

	tests := []struct {
		name    string
		reqArgs RequestArgs
		expResp ResponseTest
	}{
		{
			name: "correct token test",
			reqArgs: RequestArgs{
				Method:      "POST",
				URL:         "/auth/refresh",
				QueryParams: nil,
				Body: &requestBody{
					RefreshToken: "qwert",
				},
			},
			expResp: ResponseTest{
				Code: 200,
				Body: `{"access_token":"test_access_token","refresh_token":"test_refresh_token"}`,
			},
		},
		{
			name: "token not exist test",
			reqArgs: RequestArgs{
				Method:      "POST",
				URL:         "/auth/refresh",
				QueryParams: nil,
				Body:        &requestBody{RefreshToken: invalidTokenNotExist},
			},
			expResp: ResponseTest{
				Code: 401,
				Body: fmt.Sprintf(`{"error":"no such session"}`),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := tt.reqArgs.GetRequest()
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			strings.TrimSpace(rr.Body.String())

			assert.Equal(t, tt.expResp, ResponseTest{
				Code: rr.Code,
				Body: strings.TrimSpace(rr.Body.String()),
			})
		})
	}
}
