package main

import (
	"errors"
	"net/http"

	"github.com/kir0108/PayShareBackend/internal/data/models"
)

var ErrorInvalidCredentials = errors.New("invalid credentials")

func (app *application) loginHandler(w http.ResponseWriter, r *http.Request) {
	input := &struct {
		Api   string `json:"auth_api"`
		Token string `json:"token"`
	}{}

	if err := app.readJSON(w, r, input); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	api, err := app.api.GetAPI(input.Api)
	if err != nil {
		app.badRequestResponse(w, r, err)
	}

	var user *models.User

	if app.config.IsTest {
		user = &models.User{
			APIId:      input.Api + "test",
			APIName:    input.Api,
			FirstName:  "Test",
			SecondName: input.Api,
			ImageURL:   "",
		}
	} else {
		user, err = api.GetUser(input.Token)
		if err != nil {
			app.serverErrorResponse(w, r, err)
		}
	}


	registeredUser, err := app.users.GetByAPI(r.Context(), user.APIName, user.APIId)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNoRecord):
			if err := app.users.Add(r.Context(), user); err != nil {
				app.serverErrorResponse(w, r, err)
				return
			}
		default:
			app.serverErrorResponse(w, r, err)
			return
		}
	}

	registeredUser = user

	accessToken, err := app.jwts.CreateToken(registeredUser.Id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	refreshToken := ""
	for {
		refreshToken, err = app.generateRandomString(app.config.SessionTokenLength)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		if err := app.refreshTokens.Add(r.Context(), refreshToken, registeredUser.Id); err != nil {
			if errors.Is(err, models.ErrAlreadyExists) {
				continue
			}
			app.serverErrorResponse(w, r, err)
			return
		}

		break
	}

	var response interface{}

	response = &struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	if err := app.writeJSON(w, http.StatusOK, response, nil); err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) refreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	input := &struct {
		RefreshToken string `json:"refresh_token"`
	}{}

	if err := app.readJSON(w, r, input); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	id, err := app.refreshTokens.Get(r.Context(), input.RefreshToken)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNoRecord):
			app.unauthorizedResponse(w, r, errors.New("no such session"))
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	accessToken, err := app.jwts.CreateToken(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	newRefreshToken := ""
	for {
		newRefreshToken, err = app.generateRandomString(app.config.SessionTokenLength)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		if err := app.refreshTokens.Refresh(r.Context(), input.RefreshToken, newRefreshToken); err != nil {
			if errors.Is(err, models.ErrAlreadyExists) {
				continue
			}
			app.serverErrorResponse(w, r, err)
			return
		}

		break
	}

	response := &struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}

	if err := app.writeJSON(w, http.StatusOK, response, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) logoutUserHandler(w http.ResponseWriter, r *http.Request) {
	input := &struct {
		RefreshToken string `json:"refresh_token"`
	}{}

	if err := app.readJSON(w, r, input); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := app.refreshTokens.Delete(r.Context(), input.RefreshToken); err != nil {
		switch {
		case errors.Is(err, models.ErrNoRecord):
			app.unauthorizedResponse(w, r, errors.New("no such session"))
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}
