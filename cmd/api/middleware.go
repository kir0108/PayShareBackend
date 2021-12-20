package main

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"strings"

	"github.com/kir0108/PayShareBackend/internal/jwt"
)

type contextKey string

const (
	contextKeyID        = contextKey("id")
	contextKeyUser      = contextKey("user")
	contextKeyRoomId    = contextKey("room_id")
	contextKeyHelp      = contextKey("help")
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

func (app *application) roomIdCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		roomId, err := strconv.ParseInt(chi.URLParam(r, "room_id"), 10, 64)
		if err != nil {
			app.notFoundResponse(w, r)
			return
		}

		roomIdCtx := context.WithValue(r.Context(), contextKeyRoomId, roomId)
		next.ServeHTTP(w, r.WithContext(roomIdCtx))
	})
}

func (app *application) isRoomOwner(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, ok := r.Context().Value(contextKeyID).(int64)
		if !ok {
			app.serverErrorResponse(w, r, ErrCantRetrieveID)
			return
		}

		roomId, ok := r.Context().Value(contextKeyRoomId).(int64)
		if !ok {
			app.serverErrorResponse(w, r, ErrCantRetrieveID)
			return
		}

		room, err := app.rooms.GetById(r.Context(), roomId)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		if room.OwnerId == id {
			next.ServeHTTP(w, r)
		} else {
			app.badRequestResponse(w, r, errors.New("is not owner"))
		}

		return
	})
}

func (app *application) isHelp(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		help, err := strconv.ParseBool(r.URL.Query().Get("help"))
		if err != nil {
			help = false
		}

		helpCtx := context.WithValue(r.Context(), contextKeyHelp, help)
		next.ServeHTTP(w, r.WithContext(helpCtx))
	})
}
