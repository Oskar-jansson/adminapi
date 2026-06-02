package services

import (
	"context"
	"fmt"

	"github.com/Oskar-jansson/adminapi/models"
)

type AdministratorService struct {
	sc clientInterface
}

func (c *AdministratorService) Get(ctx context.Context, id int, params map[string]string) (*models.Administrator, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", fmt.Sprintf("/administrator/%d", id)).
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.Administrator](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *AdministratorService) List(ctx context.Context, params map[string]string) (*models.AdministratorList, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", "/administrator").
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.AdministratorList](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
