package services

import (
	"context"
	"fmt"

	"github.com/Oskar-jansson/adminapi/models"
)

type ConnectionService struct {
	sc clientInterface
}

func (c *ConnectionService) Get(ctx context.Context, id int, params map[string]string) (*models.Connection, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", fmt.Sprintf("/connection/%d", id)).
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.Connection](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *ConnectionService) List(ctx context.Context, params map[string]string) (*models.ConnectionList, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", "/connection").
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.ConnectionList](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
