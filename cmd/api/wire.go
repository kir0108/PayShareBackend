//+build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/kir0108/PayShareBackend/internal/data/postgres"
	"github.com/kir0108/PayShareBackend/internal/data/redis"
	"github.com/kir0108/PayShareBackend/internal/jwt"
)

func initApp() (*application, func(), error) {
	wire.Build(
		getConfig,
		newLogger,
		newPostgresConfig,
		postgres.NewPsqlPool,
		wire.Struct(new(postgres.UserRepo), "*"),
		wire.Struct(new(postgres.RoomRepo), "*"),
		wire.Struct(new(postgres.ParticipantRepo), "*"),
		wire.Struct(new(postgres.PurchaseRepo), "*"),
		newJwtConfig,
		jwt.NewManger,
		newRedisConfig,
		redis.NewRedisPool,
		redis.NewRefreshTokenRepository,
		redis.NewCodesRepository,
		newAPI,
		wire.Struct(new(application), "*"),
	)

	return nil, nil, nil
}
