package services

import (
	"context"
	"fmt"

	"github.com/Oskar-jansson/adminapi/models"
)

type UnitService struct {
	sc clientInterface
}

func (c *UnitService) Get(ctx context.Context, id int, params map[string]string) (*models.Unit, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", fmt.Sprintf("/unit/%d", id)).
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.Unit](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *UnitService) List(ctx context.Context, params map[string]string) (*models.UnitList, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", "/unit").
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.UnitList](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
