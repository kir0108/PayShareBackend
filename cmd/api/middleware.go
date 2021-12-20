package main

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/kir0108/PayShareBackend/internal/data/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/kir0108/PayShareBackend/internal/jwt"
)

type contextKey string

const (
	contextKeyID            = contextKey("id")
	contextKeyUser          = contextKey("user")
	contextKeyRoomId        = contextKey("room_id")
	contextKeyParticipantId = contextKey("participant_id")
	contextKeyPurchaseId    = contextKey("purchase_id")
	contextKeyHelp          = contextKey("help")
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

		if _, err := app.rooms.GetById(r.Context(), roomId); err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				app.badRequestResponse(w, r, errors.New("room dont exists"))
				return
			}

			app.serverErrorResponse(w, r, err)
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
	})
}

func (app *application) isRoomParticipants(next http.Handler) http.Handler {
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

		exist, err := app.participants.Exist(r.Context(), id, roomId)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		if !exist {
			app.badRequestResponse(w, r, errors.New("is not participant"))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (app *application) participantIdCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		participantId, err := strconv.ParseInt(chi.URLParam(r, "participant_id"), 10, 64)
		if err != nil {
			app.notFoundResponse(w, r)
			return
		}

		participantIdCtx := context.WithValue(r.Context(), contextKeyParticipantId, participantId)
		next.ServeHTTP(w, r.WithContext(participantIdCtx))
	})
}

func (app *application) purchaseIdCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		purchaseId, err := strconv.ParseInt(chi.URLParam(r, "purchase_id"), 10, 64)
		if err != nil {
			app.notFoundResponse(w, r)
			return
		}

		purchaseIdCtx := context.WithValue(r.Context(), contextKeyPurchaseId, purchaseId)
		next.ServeHTTP(w, r.WithContext(purchaseIdCtx))
	})
}

func (app *application) isPurchaseOwner(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId, ok := r.Context().Value(contextKeyID).(int64)
		if !ok {
			app.serverErrorResponse(w, r, ErrCantRetrieveID)
			return
		}

		roomId, ok := r.Context().Value(contextKeyRoomId).(int64)
		if !ok {
			app.serverErrorResponse(w, r, ErrCantRetrieveID)
			return
		}

		purchaseId, ok := r.Context().Value(contextKeyPurchaseId).(int64)
		if !ok {
			app.serverErrorResponse(w, r, ErrCantRetrieveID)
			return
		}

		participantId, err := app.participants.GetParticipantId(r.Context(), userId, roomId)
		if !ok {
			app.serverErrorResponse(w, r, err)
			return
		}

		purchase, err := app.purchases.GetById(r.Context(), purchaseId)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		if purchase.OwnerId != participantId {
			app.badRequestResponse(w, r, errors.New("is no purchase owner"))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (app *application) roomNotClosed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		if room.Close {
			app.badRequestResponse(w, r, errors.New("room closed"))
			return
		}

		next.ServeHTTP(w, r)
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
