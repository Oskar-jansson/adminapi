package services

import (
	"context"

	"github.com/Oskar-jansson/adminapi/models"
)

type AuthService struct {
	sc clientInterface
}

func (c *AuthService) Login(ctx context.Context, credentials models.Credentials) error {
	resp, err := NewRequest(ctx, c.sc).
		Method("POST", "/login").
		Body(credentials).
		Do()

	if err != nil {
		return err
	}

	// expected response
	type loginResponse struct {
		AccessToken string `json:"accesstoken"`
	}

	obj, err := convertToStruct[loginResponse](resp)
	if err != nil {
		return err
	}

	c.sc.SetToken(obj.AccessToken)
	return err
}

func (c *AuthService) Logout(ctx context.Context) error {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", "/logout").
		WithBearer().
		Do()
	if err != nil {
		return err
	}

	_, err = convertToStruct[struct{}](resp)
	return err
}
