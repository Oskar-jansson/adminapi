package services

import (
	"context"
	"fmt"

	"github.com/Oskar-jansson/adminapi/models"
)

type DomainService struct {
	sc clientInterface
}

func (c *DomainService) Get(ctx context.Context, id int, params map[string]string) (*models.Domain, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", fmt.Sprintf("/domain/%d", id)).
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.Domain](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *DomainService) List(ctx context.Context, params map[string]string) (*models.DomainList, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", "/domain").
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.DomainList](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
