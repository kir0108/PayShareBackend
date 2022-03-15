package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kir0108/PayShareBackend/internal/auth_api"
	"github.com/kir0108/PayShareBackend/internal/data/models"
	"github.com/kir0108/PayShareBackend/internal/data/postgres"
	"github.com/kir0108/PayShareBackend/internal/data/redis"
	"github.com/kir0108/PayShareBackend/internal/jwt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func getRoomIdContext(ctx context.Context, roomId int64) context.Context {
	roomIdCtx := context.WithValue(ctx, contextKeyRoomId, roomId)
	return roomIdCtx
}

func getParticipantIdContext(ctx context.Context, participantId int64) context.Context {
	participantIdCtx := context.WithValue(ctx, contextKeyParticipantId, participantId)
	return participantIdCtx
}

func Test_application_createRoomHandler(t *testing.T) {
	t.Parallel()
	app := GetTestApplication()
	handler := http.HandlerFunc(app.createRoomHandler)

	type createRoomBody struct {
		RoomName string `json:"room_name"`
		RoomDate string `json:"room_date"`
	}

	tests := []struct {
		name    string
		reqArgs RequestArgs
		expResp ResponseTest
		userId  int64
	}{
		{
			name: "create room success test",
			reqArgs: RequestArgs{
				Method: "POST",
				URL:    "/user/room",
				Body: createRoomBody{
					RoomName: "room name",
					RoomDate: "01.08.2100",
				},
			},
			expResp: ResponseTest{
				Code: 200,
				Body: fmt.Sprintf(`{"id":%d}`, DefaultRoomId),
			},
			userId: DefaultUserId,
		},
		{
			name: "create room error test",
			reqArgs: RequestArgs{
				Method: "POST",
				URL:    "/user/room",
				Body: createRoomBody{
					RoomName: ErrOtherRoomByName,
					RoomDate: "01.08.2100",
				},
			},
			expResp: ResponseTest{
				Code: 500,
				Body: `{"error":"the server encountered a problem and could not process your request"}`,
			},
			userId: DefaultUserId,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := tt.reqArgs.GetRequest()
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			userIdCtx := getUserIdContext(context.Background(), tt.userId)
			handler.ServeHTTP(rr, req.WithContext(userIdCtx))

			assert.Equal(t, tt.expResp, ResponseTest{
				Code: rr.Code,
				Body: strings.TrimSpace(rr.Body.String()),
			})
		})
	}
}

func Test_application_deleteRoomHandler(t *testing.T) {
	t.Parallel()
	app := GetTestApplication()
	handler := http.HandlerFunc(app.deleteRoomHandler)

	tests := []struct {
		name    string
		reqArgs RequestArgs
		expResp ResponseTest
		user    models.User
		roomId  int64
	}{
		{
			name: "delete room success test",
			reqArgs: RequestArgs{
				Method: "DELETE",
				URL:    fmt.Sprintf("/user/room/%d", DefaultRoomId),
			},
			expResp: ResponseTest{
				Code: 200,
			},
			user: models.User{
				Id:         DefaultUserId,
				APIId:      "api_id",
				APIName:    "google",
				FirstName:  "Daniil",
				SecondName: "Firsov",
				ImageURL:   "",
			},
			roomId: DefaultRoomId,
		},
		{
			name: "delete room not owner test",
			reqArgs: RequestArgs{
				Method: "DELETE",
				URL:    fmt.Sprintf("/user/room/%d", DefaultRoomId),
			},
			expResp: ResponseTest{
				Code: 400,
				Body: `{"error":"user is not owner"}`,
			},
			user: models.User{
				Id:         DefaultUserId + 1,
				APIId:      "api_id",
				APIName:    "google",
				FirstName:  "Daniil",
				SecondName: "Firsov",
				ImageURL:   "",
			},
			roomId: DefaultRoomId,
		},
		{
			name: "delete room not exist test",
			reqArgs: RequestArgs{
				Method: "DELETE",
				URL:    fmt.Sprintf("/user/room/%d", ErrNoExistRoomById),
			},
			expResp: ResponseTest{
				Code: 404,
				Body: `{"error":"the requested resource could not be found"}`,
			},
			user: models.User{
				Id:         DefaultUserId,
				APIId:      "api_id",
				APIName:    "google",
				FirstName:  "Daniil",
				SecondName: "Firsov",
				ImageURL:   "",
			},
			roomId: ErrNoExistRoomById,
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
			roomIdCtx := getRoomIdContext(userCtx, tt.roomId)
			handler.ServeHTTP(rr, req.WithContext(roomIdCtx))

			assert.Equal(t, tt.expResp, ResponseTest{
				Code: rr.Code,
				Body: strings.TrimSpace(rr.Body.String()),
			})
		})
	}
}

func Test_application_deleteRoomParticipantHandler(t *testing.T) {
	t.Parallel()
	app := GetTestApplication()
	handler := http.HandlerFunc(app.deleteRoomParticipantHandler)

	tests := []struct {
		name          string
		reqArgs       RequestArgs
		expResp       ResponseTest
		participantId int64
	}{
		{
			name: "delete room participant success test",
			reqArgs: RequestArgs{
				Method: "DELETE",
				URL:    "/user/room/1/participant/" + strconv.FormatInt(DefaultParticipantId, 10),
			},
			expResp: ResponseTest{
				Code: 200,
			},
			participantId: DefaultParticipantId,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := tt.reqArgs.GetRequest()
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			ctx := getParticipantIdContext(context.Background(), tt.participantId)
			handler.ServeHTTP(rr, req.WithContext(ctx))

			assert.Equal(t, tt.expResp, ResponseTest{
				Code: rr.Code,
			})
		})
	}
}

func Test_application_getClosedRoomsListHandler(t *testing.T) {
	t.Parallel()
	app := GetTestApplication()
	handler := http.HandlerFunc(app.getClosedRoomsListHandler)

	responseTest1 := &struct {
		Rooms []*models.RoomElement `json:"rooms"`
	}{
		Rooms: []*models.RoomElement{
			{
				Room: &models.Room{
					Id:       1,
					OwnerId:  DefaultUserId,
					RoomName: "test",
					RoomDate: "test",
					Close:    true,
				},
				Purchases: []*models.Purchase{
					{
						Id:      1,
						OwnerId: DefaultUserId,
						RoomId:  1,
						PName:   "test1",
						Locate: &models.Locate{
							Lat:         12,
							Long:        13,
							ShopName:    "Shop1",
							Date:        "01.02.03",
							Description: "test",
						},
						Cost: 1000,
					},
					{
						Id:      2,
						OwnerId: 2,
						RoomId:  1,
						PName:   "test2",
						Locate:  nil,
						Cost:    700,
					},
				},
				IsYour: true,
			},
			{
				Room: &models.Room{
					Id:       2,
					OwnerId:  3,
					RoomName: "test",
					RoomDate: "test",
					Close:    true,
				},
				Purchases: nil,
				IsYour:    false,
			},
		},
	}

	strJson, err := json.Marshal(responseTest1)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		reqArgs RequestArgs
		expResp ResponseTest
		userId  int64
	}{
		{
			name: "get closed rooms test",
			reqArgs: RequestArgs{
				Method: "GET",
				URL:    "/user/room/closed",
			},
			expResp: ResponseTest{
				Code: 200,
				Body: strings.TrimSpace(string(strJson)),
			},
			userId: DefaultUserId,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := tt.reqArgs.GetRequest()
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			ctx := getUserIdContext(context.Background(), tt.userId)
			handler.ServeHTTP(rr, req.WithContext(ctx))

			actualBody := strings.TrimSpace(rr.Body.String())

			assert.Equal(t, tt.expResp, ResponseTest{
				Code: rr.Code,
				Body: actualBody,
			})
		})
	}
}

func Test_application_getOpenedRoomsListHandler(t *testing.T) {
	t.Parallel()
	app := GetTestApplication()
	handler := http.HandlerFunc(app.getOpenedRoomsListHandler)

	responseTest1 := &struct {
		Rooms []*models.RoomElement `json:"rooms"`
	}{
		Rooms: []*models.RoomElement{
			{
				Room: &models.Room{
					Id:       1,
					OwnerId:  DefaultUserId,
					RoomName: "test",
					RoomDate: "test",
					Close:    false,
				},
				Purchases: []*models.Purchase{
					{
						Id:      1,
						OwnerId: DefaultUserId,
						RoomId:  1,
						PName:   "test1",
						Locate: &models.Locate{
							Lat:         12,
							Long:        13,
							ShopName:    "Shop1",
							Date:        "01.02.03",
							Description: "test",
						},
						Cost: 1000,
					},
					{
						Id:      2,
						OwnerId: 2,
						RoomId:  1,
						PName:   "test2",
						Locate:  nil,
						Cost:    700,
					},
				},
				IsYour: true,
			},
			{
				Room: &models.Room{
					Id:       2,
					OwnerId:  3,
					RoomName: "test",
					RoomDate: "test",
					Close:    false,
				},
				Purchases: nil,
				IsYour:    false,
			},
		},
	}

	strJson, err := json.Marshal(responseTest1)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		reqArgs RequestArgs
		expResp ResponseTest
		userId  int64
	}{
		{
			name: "get opened rooms test",
			reqArgs: RequestArgs{
				Method: "GET",
				URL:    "/user/room/opened",
			},
			expResp: ResponseTest{
				Code: 200,
				Body: strings.TrimSpace(string(strJson)),
			},
			userId: DefaultUserId,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := tt.reqArgs.GetRequest()
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			ctx := getUserIdContext(context.Background(), tt.userId)
			handler.ServeHTTP(rr, req.WithContext(ctx))

			actualBody := strings.TrimSpace(rr.Body.String())

			assert.Equal(t, tt.expResp, ResponseTest{
				Code: rr.Code,
				Body: actualBody,
			})
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
