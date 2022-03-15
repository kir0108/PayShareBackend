package main

import (
	"context"
	"errors"
	"github.com/kir0108/PayShareBackend/internal/data/models"
)

const (
	ErrAlreadyExistUserName = "Anton"
	ErrOtherUserName        = "Grisha"
	ErrAlreadyExistApiName  = "google"
	ErrAlreadyExistApiId    = "hcgdksbducfb"
	ErrNoRecordById         = 12
	ErrOtherById            = 13
	ErrOtherByApiId         = "user_repo_err_other_api_id"
)

type MockUsersRepo struct{}

func (m MockUsersRepo) Add(ctx context.Context, user *models.User) error {
	switch {
	case user.APIName == ErrAlreadyExistApiName && user.APIId == ErrAlreadyExistApiId:
		return models.ErrAlreadyExists
	case user.FirstName == ErrOtherUserName:
		return errors.New("other error")
	default:
		return nil
	}
}

func (m MockUsersRepo) Update(ctx context.Context, user *models.User) error {
	switch {
	case user.APIName == ErrAlreadyExistApiName && user.APIId == ErrAlreadyExistApiId:
		return models.ErrAlreadyExists
	case user.FirstName == ErrOtherUserName:
		return errors.New("other error")
	default:
		return nil
	}
}

func (m MockUsersRepo) Delete(ctx context.Context, id int64) error {
	switch id {
	case ErrNoRecordById:
		return models.ErrNoRecord
	case ErrOtherById:
		return errors.New("other error")
	default:
		return nil
	}
}

func (m MockUsersRepo) GetById(ctx context.Context, id int64) (*models.User, error) {
	switch id {
	case ErrNoRecordById:
		return nil, models.ErrNoRecord
	case ErrOtherById:
		return nil, errors.New("other error")
	default:
		return &models.User{
			Id:         id,
			APIId:      CorrectTestApiId,
			APIName:    GoogleApiName,
			FirstName:  "Daniil",
			SecondName: "Firsov",
			ImageURL:   "",
		}, nil
	}
}

func (m MockUsersRepo) GetByAPI(ctx context.Context, apiId string, apiName string) (*models.User, error) {
	switch {
	case apiId == IncorrectTestApiId || (apiName != GoogleApiName && apiName != VkApiName):
		return nil, models.ErrNoRecord
	case apiId == ErrOtherByApiId:
		return nil, errors.New("other error")
	default:
		return &models.User{
			Id:         DefaultUserId,
			APIId:      apiId,
			APIName:    apiName,
			FirstName:  "Daniil",
			SecondName: "Firsov",
			ImageURL:   "",
		}, nil
	}
}
