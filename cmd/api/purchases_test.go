package main

import (
	"github.com/kir0108/PayShareBackend/internal/auth_api"
	"github.com/kir0108/PayShareBackend/internal/data/postgres"
	"github.com/kir0108/PayShareBackend/internal/data/redis"
	"github.com/kir0108/PayShareBackend/internal/jwt"
	"go.uber.org/zap"
	"net/http"
	"testing"
)

func Test_application_addPurchaseHandler(t *testing.T) {
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
			app.addPurchaseHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_application_deletePurchaseHandler(t *testing.T) {
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
			app.deletePurchaseHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_application_joinOrLeaveToPurchaseHandler(t *testing.T) {
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
			app.joinOrLeaveToPurchaseHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_application_setPaidPurchaseParamHandler(t *testing.T) {
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
			app.setPaidPurchaseParamHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_application_updatePurchaseHandler(t *testing.T) {
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
			app.updatePurchaseHandler(tt.args.w, tt.args.r)
		})
	}
}
