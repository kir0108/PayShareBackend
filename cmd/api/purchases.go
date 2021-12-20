package main

import (
	"net/http"

	"github.com/kir0108/PayShareBackend/internal/data/models"
)

func (app *application) addPurchaseHandler(w http.ResponseWriter, r *http.Request) {
	help, ok := r.Context().Value(contextKeyHelp).(bool)
	if help && ok {
		resp := app.getHelpResponse(models.Purchase{}, struct {
			Id int64 `json:"id"`
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

	userId, ok := r.Context().Value(contextKeyID).(int64)
	if !ok {
		app.serverErrorResponse(w, r, ErrCantRetrieveID)
		return
	}

	participantId, err := app.participants.GetParticipantId(r.Context(), userId, roomId)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	purchase := &models.Purchase{}

	if err := app.readJSON(w, r, purchase); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	purchase.RoomId = roomId
	purchase.OwnerId = participantId

	if err := app.purchases.Add(r.Context(), purchase); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	response := &struct {
		Id int64 `json:"id"`
	}{
		Id: purchase.Id,
	}

	if err := app.writeJSON(w, http.StatusOK, response, nil); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) updatePurchaseHandler(w http.ResponseWriter, r *http.Request) {
	help, ok := r.Context().Value(contextKeyHelp).(bool)
	if help && ok {
		resp := app.getHelpResponse(struct {
			Name   string         `json:"name"`
			Locate *models.Locate `json:"locate"`
			Cost   int64          `json:"cost"`
		}{}, nil)

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

	participantId, ok := r.Context().Value(contextKeyParticipantId).(int64)
	if !ok {
		app.serverErrorResponse(w, r, ErrCantRetrieveID)
		return
	}

	purchase := &models.Purchase{}

	if err := app.readJSON(w, r, purchase); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	purchase.RoomId = roomId
	purchase.OwnerId = participantId

	if err := app.purchases.Add(r.Context(), purchase); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	response := &struct {
		Id int64 `json:"id"`
	}{
		Id: purchase.Id,
	}

	if err := app.writeJSON(w, http.StatusOK, response, nil); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) deletePurchaseHandler(w http.ResponseWriter, r *http.Request) {
	help, ok := r.Context().Value(contextKeyHelp).(bool)
	if help && ok {
		resp := app.getHelpResponse(nil, nil)

		if err := app.writeJSON(w, http.StatusOK, resp, nil); err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		return
	}

	purchaseId, ok := r.Context().Value(contextKeyPurchaseId).(int64)
	if !ok {
		app.serverErrorResponse(w, r, ErrCantRetrieveID)
		return
	}

	if err := app.purchases.Delete(r.Context(), purchaseId); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
