package main

import (
	"errors"
	"github.com/kir0108/PayShareBackend/internal/auth_api"
	"github.com/kir0108/PayShareBackend/internal/data/postgres"
	"github.com/kir0108/PayShareBackend/internal/data/redis"
	"github.com/kir0108/PayShareBackend/internal/jwt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func Test_application_badRequestResponse(t *testing.T) {
	t.Parallel()
	app := GetTestApplication()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.badRequestResponse(w, r, errors.New("test"))
	})

	tests := []struct {
		name    string
		reqArgs RequestArgs
		expResp ResponseTest
	}{
		{
			name: "bad request response test",
			reqArgs: RequestArgs{
				Method:      "POST",
				URL:         "/auth/refresh",
				QueryParams: nil,
			},
			expResp: ResponseTest{
				Code: 400,
				Body: `{"error":"test"}`,
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

func Test_application_clientErrorResponse(t *testing.T) {
	t.Parallel()
	app := GetTestApplication()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.clientErrorResponse(w, r, http.StatusForbidden, "test")
	})

	tests := []struct {
		name    string
		reqArgs RequestArgs
		expResp ResponseTest
	}{
		{
			name: "client error response test",
			reqArgs: RequestArgs{
				Method:      "POST",
				URL:         "/auth/refresh",
				QueryParams: nil,
			},
			expResp: ResponseTest{
				Code: 403,
				Body: `{"error":"test"}`,
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

func Test_application_errorResponse(t *testing.T) {
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
		w       http.ResponseWriter
		r       *http.Request
		status  int
		mesaage interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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
			app.errorResponse(tt.args.w, tt.args.r, tt.args.status, tt.args.mesaage)
		})
	}
}

func Test_application_failedValidationResponse(t *testing.T) {
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
		w      http.ResponseWriter
		r      *http.Request
		errors map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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
			app.failedValidationResponse(tt.args.w, tt.args.r, tt.args.errors)
		})
	}
}

func Test_application_forbiddenResponse(t *testing.T) {
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
		w        http.ResponseWriter
		r        *http.Request
		messsage string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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
			app.forbiddenResponse(tt.args.w, tt.args.r, tt.args.messsage)
		})
	}
}

func Test_application_logError(t *testing.T) {
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
		r   *http.Request
		err error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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
			app.logError(tt.args.r, tt.args.err)
		})
	}
}

func Test_application_methodNotAllowedResponse(t *testing.T) {
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
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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
			app.methodNotAllowedResponse(tt.args.w, tt.args.r)
		})
	}
}

func Test_application_notFoundResponse(t *testing.T) {
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
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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
			app.notFoundResponse(tt.args.w, tt.args.r)
		})
	}
}

func Test_application_userNoOwnerResponse(t *testing.T) {
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
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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
			app.userNoOwnerResponse(tt.args.w, tt.args.r)
		})
	}
}

func Test_application_userAlreadyExistsResponse(t *testing.T) {
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
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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
			app.userAlreadyExistsResponse(tt.args.w, tt.args.r)
		})
	}
}

func Test_application_unauthorizedResponse(t *testing.T) {
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
		err error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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
			app.unauthorizedResponse(tt.args.w, tt.args.r, tt.args.err)
		})
	}
}

func Test_application_serverErrorResponse(t *testing.T) {
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
		err error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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
			app.serverErrorResponse(tt.args.w, tt.args.r, tt.args.err)
		})
	}
}

func Test_application_roomNotClosed(t *testing.T) {
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
		next http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   http.Handler
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
			if got := app.roomNotClosed(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("roomNotClosed() = %v, want %v", got, tt.want)
			}
		})
	}
}
