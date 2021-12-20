package main

import (
	"net/http"

	"github.com/kir0108/PayShareBackend/internal/data/models"
)

func (app *application) getUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	help, ok := r.Context().Value(contextKeyHelp).(bool)
	if help && ok {
		resp := app.getHelpResponse(nil, models.User{})

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

	if err := app.writeJSON(w, http.StatusOK, user, nil); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) updateUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	help, ok := r.Context().Value(contextKeyHelp).(bool)
	if help && ok {
		resp := app.getHelpResponse(models.User{}, nil)

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

	newUser := &models.User{}

	if err := app.readJSON(w, r, newUser); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	newUser.Id = user.Id
	newUser.APIId = user.APIId
	newUser.APIName = user.APIName

	if err := app.users.Update(r.Context(), newUser); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
