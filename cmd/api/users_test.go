package main

import (
	"context"
	"fmt"
	"github.com/kir0108/PayShareBackend/internal/data/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func getUserIdContext(ctx context.Context, userId int64) context.Context {
	userIdCtx := context.WithValue(ctx, contextKeyID, userId)
	return userIdCtx
}

func getUserContext(ctx context.Context, user models.User) context.Context {
	userCtx := context.WithValue(ctx, contextKeyUser, &user)
	return userCtx
}

func Test_application_getUserProfileHandler(t *testing.T) {
	t.Parallel()
	app := GetTestApplication()
	handler := http.HandlerFunc(app.getUserProfileHandler)

	tests := []struct {
		name    string
		reqArgs RequestArgs
		expResp ResponseTest
		user    models.User
	}{
		{
			name: "get user test",
			reqArgs: RequestArgs{
				Method: "GET",
				URL:    "/user",
			},
			expResp: ResponseTest{
				Code: 200,
			},
			user: models.User{
				Id:         1,
				APIId:      "jeg1y23",
				APIName:    "google",
				FirstName:  "Daniil",
				SecondName: "Firsov",
				ImageURL:   "http://fhwuqmn/img.png",
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
			userCtx := getUserContext(context.Background(), tt.user)
			fmt.Println(userCtx)
			handler.ServeHTTP(rr, req.WithContext(userCtx))

			tt.expResp.Body = fmt.Sprintf(`{"first_name":"%s","second_name":"%s","image_url":"%s"}`,
				tt.user.FirstName, tt.user.SecondName, tt.user.ImageURL)

			assert.Equal(t, tt.expResp, ResponseTest{
				Code: rr.Code,
				Body: strings.TrimSpace(rr.Body.String()),
			})
		})
	}
}

func Test_application_updateUserProfileHandler(t *testing.T) {
	t.Parallel()
	app := GetTestApplication()
	handler := http.HandlerFunc(app.updateUserProfileHandler)

	tests := []struct {
		name    string
		reqArgs RequestArgs
		expResp ResponseTest
		user    models.User
	}{
		{
			name: "update user success test",
			reqArgs: RequestArgs{
				Method: "PUT",
				URL:    "/user",
				Body: &models.User{
					FirstName:  "Keril",
					SecondName: "Korkunov",
					ImageURL:   "some_url",
				},
			},
			expResp: ResponseTest{
				Code: 200,
			},
			user: models.User{
				Id:         1,
				APIId:      "jeg1y23",
				APIName:    "google",
				FirstName:  "Daniil",
				SecondName: "Firsov",
				ImageURL:   "http://fhwuqmn/img.png",
			},
		},
		{
			name: "update user server error test",
			reqArgs: RequestArgs{
				Method: "PUT",
				URL:    "/user",
				Body: &models.User{
					FirstName:  "Grisha",
					SecondName: "Korkunov",
					ImageURL:   "some_url",
				},
			},
			expResp: ResponseTest{
				Code: 500,
			},
			user: models.User{
				Id:         1,
				APIId:      "jeg1y23",
				APIName:    "google",
				FirstName:  "Daniil",
				SecondName: "Firsov",
				ImageURL:   "http://fhwuqmn/img.png",
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
			userCtx := getUserContext(context.Background(), tt.user)
			handler.ServeHTTP(rr, req.WithContext(userCtx))

			assert.Equal(t, tt.expResp, ResponseTest{
				Code: rr.Code,
			})
		})
	}
}
