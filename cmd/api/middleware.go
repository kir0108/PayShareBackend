package main

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/kir0108/PayShareBackend/internal/jwt"
)

type contextKey string

const (
	contextKeyID   = contextKey("id")
	contextKeyUser = contextKey("user")
)

var ErrCantRetrieveID = errors.New("can't retrieve id")

func (app *application) auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			app.unauthorizedResponse(w, r, errors.New("no token provided"))
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")

		id, err := app.jwts.GetIdFromToken(token)
		if err != nil {
			invalidTokenErr := &jwt.InvalidTokenError{}
			switch {
			case errors.As(err, &invalidTokenErr):
				app.unauthorizedResponse(w, r, invalidTokenErr)
			default:
				app.serverErrorResponse(w, r, err)
			}
			return
		}

		idContext := context.WithValue(r.Context(), contextKeyID, id)
		next.ServeHTTP(w, r.WithContext(idContext))
	})
}

func (app *application) userCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, ok := r.Context().Value(contextKeyID).(int64)

		if !ok {
			app.serverErrorResponse(w, r, ErrCantRetrieveID)
			return
		}

		user, err := app.users.GetById(r.Context(), id)

		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		userCtx := context.WithValue(r.Context(), contextKeyUser, user)
		next.ServeHTTP(w, r.WithContext(userCtx))
	})
}
