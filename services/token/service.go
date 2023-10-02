package token

import (
	"Book_Management_System/models"
	"context"
	"encoding/json"
	"errors"
	"google.golang.org/api/idtoken"
)

type service struct {
	webClientID string
	validator   Validator
}
type Validator func(authToken, webClientID string) (*models.Users, error)

// nolint - unexported return type for exported function
func New(webClientID string) service {
	return service{
		webClientID: webClientID,
		validator:   jwtValidator,
	}
}

func (svc service) CheckUserAuthorized(authToken string) (*models.Users, error) {
	if authToken == "" {
		return nil, errors.New("missing param: authToken")
	}

	userData, err := svc.validator(authToken, svc.webClientID)
	if err != nil {
		return nil, errors.New("token is invalid")
	}

	if !userData.Verified || userData.Error.Code == 401 {
		return nil, errors.New("token is invalid or user unauthenticated ")
	}

	return userData, nil
}

func jwtValidator(authToken, webClientID string) (*models.Users, error) {
	payload, err := idtoken.Validate(context.Background(), authToken, webClientID)
	if err != nil {
		return nil, err
	}

	return parseClaims(payload.Claims)
}

func parseClaims(claims map[string]interface{}) (*models.Users, error) {
	b, err := json.Marshal(claims)
	if err != nil {
		return nil, err
	}

	var data models.Users

	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
