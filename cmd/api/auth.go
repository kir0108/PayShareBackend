package main

import (
	"errors"
	"net/http"

	"github.com/kir0108/PayShareBackend/internal/data/models"
)

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
		return
	}

	var user *models.User

	imageUrl := map[string]string{
		"google": "https://www.m24.ru/b/d/nBkSUhL2jVUgm82xPqzJrMCqzJ3w-pun2XyQ2q2C_2OZcGuaSnvVjCdg4M4S7FjDvM_AtC_QbIk8W7zj1GdwKSGT_w=VKxfke3fPJFrcYKGoS-5yg.jpg",
		"vk":     "https://static.wikia.nocookie.net/gorillaz/images/0/04/Noooodes.jpg/revision/latest?cb=20181101202938&path-prefix=ru",
	}

	if app.config.IsTest {
		user = &models.User{
			APIId:      input.Api + "test",
			APIName:    input.Api,
			FirstName:  "Test",
			SecondName: input.Api,
			ImageURL:   imageUrl[input.Api],
		}
	} else {
		user, err = api.GetUser(input.Token)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}
	}

	registeredUser, err := app.users.GetByAPI(r.Context(), user.APIId, user.APIName)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNoRecord):
			if err := app.users.Add(r.Context(), user); err != nil {
				app.serverErrorResponse(w, r, err)
				return
			}

			registeredUser = user
		default:
			app.serverErrorResponse(w, r, err)
			return
		}
	}

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
