package services

import (
	"context"
	"fmt"

	"github.com/Oskar-jansson/adminapi/models"
)

type TimezoneService struct {
	sc clientInterface
}

func (c *TimezoneService) Get(ctx context.Context, id int, params map[string]string) (*models.Timezone, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", fmt.Sprintf("/timezone/%d", id)).
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.Timezone](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *TimezoneService) List(ctx context.Context, params map[string]string) (*models.TimezoneList, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", "/timezone").
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.TimezoneList](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
