package services

import (
	"context"

	"github.com/Oskar-jansson/adminapi/models"
)

type DateService struct {
	sc clientInterface
}

func (c *DateService) List(ctx context.Context, params map[string]string) (*models.DateList, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", "/date").
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.DateList](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
