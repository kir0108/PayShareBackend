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

func Test_application_createRoomHandler(t *testing.T) {
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
			app.createRoomHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_application_deleteRoomHandler(t *testing.T) {
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
			app.deleteRoomHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_application_deleteRoomParticipantHandler(t *testing.T) {
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
			app.deleteRoomParticipantHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_application_getClosedRoomsListHandler(t *testing.T) {
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
			app.getClosedRoomsListHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_application_getOpenedRoomsListHandler(t *testing.T) {
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
			app.getOpenedRoomsListHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_application_getRoomCodeHandler(t *testing.T) {
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
			app.getRoomCodeHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_application_getRoomHandler(t *testing.T) {
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
			app.getRoomHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_application_joinRoomHandler(t *testing.T) {
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
			app.joinRoomHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_application_joinToRoomHandler(t *testing.T) {
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
			app.joinToRoomHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_application_leaveRoomHandler(t *testing.T) {
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
			app.leaveRoomHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_application_setCloseRoomHandler(t *testing.T) {
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
			app.setCloseRoomHandler(tt.args.w, tt.args.r)
		})
	}
}
