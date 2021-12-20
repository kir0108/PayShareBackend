package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"net/http"
	"time"

	"github.com/kir0108/PayShareBackend/internal/auth_api"
	"github.com/kir0108/PayShareBackend/internal/auth_api/google_api"
	"github.com/kir0108/PayShareBackend/internal/auth_api/vk_api"
	"github.com/kir0108/PayShareBackend/internal/data/postgres"
	"github.com/kir0108/PayShareBackend/internal/data/redis"
	"github.com/kir0108/PayShareBackend/internal/jwt"
	"github.com/xlab/closer"
)

type application struct {
	config        *config
	logger        *zap.SugaredLogger
	jwts          *jwt.Manager
	codes         *redis.CodesRepository
	users         *postgres.UserRepo
	rooms         *postgres.RoomRepo
	participants  *postgres.ParticipantRepo
	purchases     *postgres.PurchaseRepo
	refreshTokens *redis.RefreshTokenRepository
	api           *auth_api.Api
}

func main() {
	app, cleanup, err := initApp()
	if err != nil {
		log.Fatal("error initiating application", err)
	}

	closer.Bind(func() {
		log.Print("Stopping server")
		cleanup()
	})

	errLogger, err := zap.NewStdLogAt(app.logger.Desugar(), zap.ErrorLevel)
	if err != nil {
		app.logger.Fatalw("error initiating server logger", "err", err)
	}

	server := &http.Server{
		Addr:         ":" + app.config.Port,
		Handler:      app.route(),
		ErrorLog:     errLogger,
		WriteTimeout: 120 * time.Second,
	}

	app.logger.Infow("Started server", "port", app.config.Port)

	app.logger.Fatalw("server error", "err", server.ListenAndServe())
}

func newLogger(c *config) (*zap.SugaredLogger, func(), error) {
	var logger *zap.Logger
	var err error

	if c.Production {
		logger, err = zap.NewProduction()
	} else {
		conf := zap.NewDevelopmentConfig()
		conf.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		logger, err = conf.Build()
	}

	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		_ = logger.Sync()
	}

	return logger.Sugar(), cleanup, nil
}

func newPostgresConfig(c *config) *postgres.Config {
	return &postgres.Config{
		PostgresUrl: c.PostgresUrl,
	}
}

func newJwtConfig(c *config) *jwt.Config {
	return &jwt.Config{
		Secret:     c.Secret,
		Expiration: c.JwtTTL,
	}
}

func newRedisConfig(c *config) *redis.Config {
	return &redis.Config{
		RedisUrl:             c.RedisUrl,
		CodeHideTTL:          c.CodeHideTTL,
		CodeExpiredTTL:       c.CodeExpiredTTL,
		SessionTTl:           c.SessionTTl,
		SessionCleanupPeriod: c.SessionCleanupPeriod,
		SessionWindowPeriod:  c.SessionWindowPeriod,
	}
}

func newVKApi(c *config) *vk_api.Config {
	return &vk_api.Config{
	}
}

func newGoogleApi(c *config) *google_api.Config {
	return &google_api.Config{
	}
}

func newAPI(c *config) *auth_api.Api {
	apis := make([]auth_api.AuthApi, 0)

	apis = append(apis, vk_api.NewVKApi(newVKApi(c)))
	apis = append(apis, google_api.NewGoogleApi(newGoogleApi(c)))

	return auth_api.NewApi(apis...)
}
