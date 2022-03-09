package main

import (
	"github.com/kir0108/PayShareBackend/internal/auth_api"
	"github.com/kir0108/PayShareBackend/internal/auth_api/google_api"
	"github.com/kir0108/PayShareBackend/internal/auth_api/vk_api"
	"github.com/kir0108/PayShareBackend/internal/data/postgres"
	"github.com/kir0108/PayShareBackend/internal/data/redis"
	"github.com/kir0108/PayShareBackend/internal/jwt"
	"go.uber.org/zap"
	"reflect"
	"testing"
)

func Test_newAPI(t *testing.T) {
	type args struct {
		c *config
	}
	tests := []struct {
		name string
		args args
		want *auth_api.Api
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newAPI(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newAPI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newGoogleApi(t *testing.T) {
	type args struct {
		c *config
	}
	tests := []struct {
		name string
		args args
		want *google_api.Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newGoogleApi(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newGoogleApi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newJwtConfig(t *testing.T) {
	type args struct {
		c *config
	}
	tests := []struct {
		name string
		args args
		want *jwt.Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newJwtConfig(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newJwtConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newLogger(t *testing.T) {
	type args struct {
		c *config
	}
	tests := []struct {
		name    string
		args    args
		want    *zap.SugaredLogger
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := newLogger(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("newLogger() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newLogger() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newPostgresConfig(t *testing.T) {
	type args struct {
		c *config
	}
	tests := []struct {
		name string
		args args
		want *postgres.Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newPostgresConfig(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newPostgresConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newRedisConfig(t *testing.T) {
	type args struct {
		c *config
	}
	tests := []struct {
		name string
		args args
		want *redis.Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newRedisConfig(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newRedisConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newVKApi(t *testing.T) {
	type args struct {
		c *config
	}
	tests := []struct {
		name string
		args args
		want *vk_api.Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newVKApi(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newVKApi() = %v, want %v", got, tt.want)
			}
		})
	}
}
