package main

import (
	"errors"
	"fmt"
	"github.com/kir0108/PayShareBackend/internal/data/models"
	"strings"
)

type MockApi struct{}

func TokenCreate(user models.User) string {
	return fmt.Sprintf("%s_%s_%s_%s", user.APIId, user.APIName, user.FirstName, user.SecondName)
}

func (m MockApi) GetUser(token string) (*models.User, error) {
	data := strings.Split(token, "_")

	if len(data) < 4 {
		return nil, errors.New("invalid test token")
	}

	return &models.User{
		APIId:      data[0],
		APIName:    data[1],
		FirstName:  data[2],
		SecondName: data[3],
	}, nil
}

func (m MockApi) GetName() string {
	return "test"
}
