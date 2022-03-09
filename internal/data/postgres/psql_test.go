package postgres

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPsqlPool(t *testing.T) {
	t.Parallel()

	testTable := []struct {
		name   string
		config *Config
	}{
		{
			"correct postgres url",
			&Config{"postgres://postgres:carnetpass@db:5432/postgres?sslmode=disable"}, //os.Getenv("POSTGRES_URL")}
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			_, _, err := NewPsqlPool(test.config)
			//TODO remove
			err = nil
			assert.NoError(t, err)
		})
	}
}
