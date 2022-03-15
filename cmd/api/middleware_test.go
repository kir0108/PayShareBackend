package main

//
//import (
//	"github.com/kir0108/PayShareBackend/internal/auth_api"
//	"github.com/kir0108/PayShareBackend/internal/data/postgres"
//	"github.com/kir0108/PayShareBackend/internal/data/redis"
//	"github.com/kir0108/PayShareBackend/internal/jwt"
//	"go.uber.org/zap"
//	"net/http"
//	"reflect"
//	"testing"
//)
//
//func Test_application_auth(t *testing.T) {
//	type fields struct {
//		config        *config
//		logger        *zap.SugaredLogger
//		jwts          *jwt.Manager
//		codes         redis.CodesRepositoryType
//		users         postgres.UserRepoType
//		rooms         postgres.RoomRepoType
//		participants  postgres.ParticipantRepoType
//		purchases     postgres.PurchaseRepoType
//		refreshTokens redis.RefreshTokenRepositoryType
//		api           *auth_api.Api
//	}
//	type args struct {
//		next http.Handler
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   http.Handler
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			app := &application{
//				config:        tt.fields.config,
//				logger:        tt.fields.logger,
//				jwts:          tt.fields.jwts,
//				codes:         tt.fields.codes,
//				users:         tt.fields.users,
//				rooms:         tt.fields.rooms,
//				participants:  tt.fields.participants,
//				purchases:     tt.fields.purchases,
//				refreshTokens: tt.fields.refreshTokens,
//				api:           tt.fields.api,
//			}
//			if got := app.auth(tt.args.next); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("auth() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_application_isPurchaseOwner(t *testing.T) {
//	type fields struct {
//		config        *config
//		logger        *zap.SugaredLogger
//		jwts          *jwt.Manager
//		codes         redis.CodesRepositoryType
//		users         postgres.UserRepoType
//		rooms         postgres.RoomRepoType
//		participants  postgres.ParticipantRepoType
//		purchases     postgres.PurchaseRepoType
//		refreshTokens redis.RefreshTokenRepositoryType
//		api           *auth_api.Api
//	}
//	type args struct {
//		next http.Handler
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   http.Handler
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			app := &application{
//				config:        tt.fields.config,
//				logger:        tt.fields.logger,
//				jwts:          tt.fields.jwts,
//				codes:         tt.fields.codes,
//				users:         tt.fields.users,
//				rooms:         tt.fields.rooms,
//				participants:  tt.fields.participants,
//				purchases:     tt.fields.purchases,
//				refreshTokens: tt.fields.refreshTokens,
//				api:           tt.fields.api,
//			}
//			if got := app.isPurchaseOwner(tt.args.next); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("isPurchaseOwner() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_application_isRoomOwner(t *testing.T) {
//	type fields struct {
//		config        *config
//		logger        *zap.SugaredLogger
//		jwts          *jwt.Manager
//		codes         redis.CodesRepositoryType
//		users         postgres.UserRepoType
//		rooms         postgres.RoomRepoType
//		participants  postgres.ParticipantRepoType
//		purchases     postgres.PurchaseRepoType
//		refreshTokens redis.RefreshTokenRepositoryType
//		api           *auth_api.Api
//	}
//	type args struct {
//		next http.Handler
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   http.Handler
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			app := &application{
//				config:        tt.fields.config,
//				logger:        tt.fields.logger,
//				jwts:          tt.fields.jwts,
//				codes:         tt.fields.codes,
//				users:         tt.fields.users,
//				rooms:         tt.fields.rooms,
//				participants:  tt.fields.participants,
//				purchases:     tt.fields.purchases,
//				refreshTokens: tt.fields.refreshTokens,
//				api:           tt.fields.api,
//			}
//			if got := app.isRoomOwner(tt.args.next); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("isRoomOwner() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_application_isRoomParticipants(t *testing.T) {
//	type fields struct {
//		config        *config
//		logger        *zap.SugaredLogger
//		jwts          *jwt.Manager
//		codes         redis.CodesRepositoryType
//		users         postgres.UserRepoType
//		rooms         postgres.RoomRepoType
//		participants  postgres.ParticipantRepoType
//		purchases     postgres.PurchaseRepoType
//		refreshTokens redis.RefreshTokenRepositoryType
//		api           *auth_api.Api
//	}
//	type args struct {
//		next http.Handler
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   http.Handler
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			app := &application{
//				config:        tt.fields.config,
//				logger:        tt.fields.logger,
//				jwts:          tt.fields.jwts,
//				codes:         tt.fields.codes,
//				users:         tt.fields.users,
//				rooms:         tt.fields.rooms,
//				participants:  tt.fields.participants,
//				purchases:     tt.fields.purchases,
//				refreshTokens: tt.fields.refreshTokens,
//				api:           tt.fields.api,
//			}
//			if got := app.isRoomParticipants(tt.args.next); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("isRoomParticipants() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_application_participantIdCtx(t *testing.T) {
//	type fields struct {
//		config        *config
//		logger        *zap.SugaredLogger
//		jwts          *jwt.Manager
//		codes         redis.CodesRepositoryType
//		users         postgres.UserRepoType
//		rooms         postgres.RoomRepoType
//		participants  postgres.ParticipantRepoType
//		purchases     postgres.PurchaseRepoType
//		refreshTokens redis.RefreshTokenRepositoryType
//		api           *auth_api.Api
//	}
//	type args struct {
//		next http.Handler
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   http.Handler
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			app := &application{
//				config:        tt.fields.config,
//				logger:        tt.fields.logger,
//				jwts:          tt.fields.jwts,
//				codes:         tt.fields.codes,
//				users:         tt.fields.users,
//				rooms:         tt.fields.rooms,
//				participants:  tt.fields.participants,
//				purchases:     tt.fields.purchases,
//				refreshTokens: tt.fields.refreshTokens,
//				api:           tt.fields.api,
//			}
//			if got := app.participantIdCtx(tt.args.next); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("participantIdCtx() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_application_purchaseIdCtx(t *testing.T) {
//	type fields struct {
//		config        *config
//		logger        *zap.SugaredLogger
//		jwts          *jwt.Manager
//		codes         redis.CodesRepositoryType
//		users         postgres.UserRepoType
//		rooms         postgres.RoomRepoType
//		participants  postgres.ParticipantRepoType
//		purchases     postgres.PurchaseRepoType
//		refreshTokens redis.RefreshTokenRepositoryType
//		api           *auth_api.Api
//	}
//	type args struct {
//		next http.Handler
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   http.Handler
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			app := &application{
//				config:        tt.fields.config,
//				logger:        tt.fields.logger,
//				jwts:          tt.fields.jwts,
//				codes:         tt.fields.codes,
//				users:         tt.fields.users,
//				rooms:         tt.fields.rooms,
//				participants:  tt.fields.participants,
//				purchases:     tt.fields.purchases,
//				refreshTokens: tt.fields.refreshTokens,
//				api:           tt.fields.api,
//			}
//			if got := app.purchaseIdCtx(tt.args.next); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("purchaseIdCtx() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_application_roomIdCtx(t *testing.T) {
//	type fields struct {
//		config        *config
//		logger        *zap.SugaredLogger
//		jwts          *jwt.Manager
//		codes         redis.CodesRepositoryType
//		users         postgres.UserRepoType
//		rooms         postgres.RoomRepoType
//		participants  postgres.ParticipantRepoType
//		purchases     postgres.PurchaseRepoType
//		refreshTokens redis.RefreshTokenRepositoryType
//		api           *auth_api.Api
//	}
//	type args struct {
//		next http.Handler
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   http.Handler
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			app := &application{
//				config:        tt.fields.config,
//				logger:        tt.fields.logger,
//				jwts:          tt.fields.jwts,
//				codes:         tt.fields.codes,
//				users:         tt.fields.users,
//				rooms:         tt.fields.rooms,
//				participants:  tt.fields.participants,
//				purchases:     tt.fields.purchases,
//				refreshTokens: tt.fields.refreshTokens,
//				api:           tt.fields.api,
//			}
//			if got := app.roomIdCtx(tt.args.next); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("roomIdCtx() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_application_userCtx(t *testing.T) {
//	type fields struct {
//		config        *config
//		logger        *zap.SugaredLogger
//		jwts          *jwt.Manager
//		codes         redis.CodesRepositoryType
//		users         postgres.UserRepoType
//		rooms         postgres.RoomRepoType
//		participants  postgres.ParticipantRepoType
//		purchases     postgres.PurchaseRepoType
//		refreshTokens redis.RefreshTokenRepositoryType
//		api           *auth_api.Api
//	}
//	type args struct {
//		next http.Handler
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   http.Handler
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			app := &application{
//				config:        tt.fields.config,
//				logger:        tt.fields.logger,
//				jwts:          tt.fields.jwts,
//				codes:         tt.fields.codes,
//				users:         tt.fields.users,
//				rooms:         tt.fields.rooms,
//				participants:  tt.fields.participants,
//				purchases:     tt.fields.purchases,
//				refreshTokens: tt.fields.refreshTokens,
//				api:           tt.fields.api,
//			}
//			if got := app.userCtx(tt.args.next); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("userCtx() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
