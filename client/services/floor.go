package services

import (
	"context"
	"fmt"

	"github.com/Oskar-jansson/adminapi/models"
)

type FloorService struct {
	sc clientInterface
}

func (c *FloorService) Get(ctx context.Context, id int, params map[string]string) (*models.Floor, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", fmt.Sprintf("/floor/%d", id)).
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.Floor](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *FloorService) List(ctx context.Context, params map[string]string) (*models.FloorList, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", "/floor").
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.FloorList](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *FloorService) Edit(ctx context.Context, id int, changes models.FloorInput) (*models.Floor, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("PATCH", fmt.Sprintf("/floor/%d", id)).
		Body(changes).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.Floor](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
