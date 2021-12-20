package main

import (
	"errors"
	"github.com/kir0108/PayShareBackend/internal/data/models"
	"net/http"
)

func (app *application) createRoomHandler(w http.ResponseWriter, r *http.Request) {
	help, ok := r.Context().Value(contextKeyHelp).(bool)
	if help && ok {
		resp := app.getHelpResponse(struct {
			RoomName string `json:"room_name"`
			RoomDate string `json:"room_date"`
		}{}, struct {
			Id int64 `json:"id"`
		}{})

		if err := app.writeJSON(w, http.StatusOK, resp, nil); err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		return
	}

	userId, ok := r.Context().Value(contextKeyID).(int64)
	if !ok {
		app.serverErrorResponse(w, r, ErrCantRetrieveID)
		return
	}

	input := &models.Room{}

	if err := app.readJSON(w, r, input); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	room := &models.Room{
		OwnerId:  userId,
		RoomName: input.RoomName,
		RoomDate: input.RoomDate,
	}

	if err := app.rooms.Add(r.Context(), room); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	response := &struct {
		Id int64 `json:"id"`
	}{
		Id: room.Id,
	}

	if err := app.writeJSON(w, http.StatusOK, response, nil); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) deleteRoomHandler(w http.ResponseWriter, r *http.Request) {
	help, ok := r.Context().Value(contextKeyHelp).(bool)
	if help && ok {
		resp := app.getHelpResponse(nil, nil)

		if err := app.writeJSON(w, http.StatusOK, resp, nil); err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		return
	}

	user, ok := r.Context().Value(contextKeyUser).(*models.User)
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
		if errors.Is(err, models.ErrNoRecord) {
			app.notFoundResponse(w, r)
			return
		}

		app.serverErrorResponse(w, r, err)
		return
	}

	if room.OwnerId != user.Id {
		app.userNoOwnerResponse(w, r)
		return
	}

	if err := app.rooms.Delete(r.Context(), roomId); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (app *application) setCloseRoomHandler(w http.ResponseWriter, r *http.Request) {
	help, ok := r.Context().Value(contextKeyHelp).(bool)
	if help && ok {
		resp := app.getHelpResponse(nil, nil)

		if err := app.writeJSON(w, http.StatusOK, resp, nil); err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		return
	}

	user, ok := r.Context().Value(contextKeyUser).(*models.User)
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
		if errors.Is(err, models.ErrNoRecord) {
			app.notFoundResponse(w, r)
			return
		}

		app.serverErrorResponse(w, r, err)
		return
	}

	if room.OwnerId != user.Id {
		app.userNoOwnerResponse(w, r)
		return
	}

	if err := app.rooms.UpdateClose(r.Context(), room.Id, !room.Close); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (app *application) getRoomCodeHandler(w http.ResponseWriter, r *http.Request) {
	help, ok := r.Context().Value(contextKeyHelp).(bool)
	if help && ok {
		resp := app.getHelpResponse(nil, struct {
			Code string `json:"code"`
		}{})

		if err := app.writeJSON(w, http.StatusOK, resp, nil); err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		return
	}

	roomId, ok := r.Context().Value(contextKeyRoomId).(int64)
	if !ok {
		app.serverErrorResponse(w, r, ErrCantRetrieveID)
		return
	}

	code, err := app.codes.GetCode(r.Context(), roomId)
	if err != nil {
		if !errors.Is(err, models.ErrNoRecord) {
			app.serverErrorResponse(w, r, err)
			return
		}

		for {
			code, err = app.generateNumberString(app.config.CodeLength)
			if err != nil {
				app.serverErrorResponse(w, r, err)
				return
			}

			if err := app.codes.Add(r.Context(), code, roomId); err != nil {
				if !errors.Is(err, models.ErrAlreadyExists) {
					app.serverErrorResponse(w, r, err)
					return
				}
			} else {
				break
			}
		}
	}

	response := &struct {
		Code string `json:"code"`
	}{
		Code: code,
	}

	if err := app.writeJSON(w, http.StatusOK, response, nil); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
