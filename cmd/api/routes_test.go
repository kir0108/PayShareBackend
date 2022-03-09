package main

import (
	"github.com/kir0108/PayShareBackend/internal/auth_api"
	"github.com/kir0108/PayShareBackend/internal/data/postgres"
	"github.com/kir0108/PayShareBackend/internal/data/redis"
	"github.com/kir0108/PayShareBackend/internal/jwt"
	"go.uber.org/zap"
	"net/http"
	"reflect"
	"testing"
)

func Test_application_route(t *testing.T) {
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
	tests := []struct {
		name   string
		fields fields
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
			if got := app.route(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("route() = %v, want %v", got, tt.want)
			}
		})
	}
}
