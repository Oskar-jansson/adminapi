package services

import (
	"context"

	"github.com/Oskar-jansson/adminapi/models"
)

type SystemService struct {
	sc clientInterface
}

func (c *SystemService) List(ctx context.Context, params map[string]string) (*models.System, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", "/system").
		Query(params).
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.System](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
