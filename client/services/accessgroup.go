package services

import (
	"context"
	"fmt"

	"github.com/Oskar-jansson/adminapi/models"
)

type AccessGroupService struct {
	sc clientInterface
}

func (c *AccessGroupService) List(ctx context.Context, params map[string]string) (*models.AccessGroupList, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", "/accessgroup").
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	return convertToStruct[models.AccessGroupList](resp)
}

func (c *AccessGroupService) Get(ctx context.Context, id int, params map[string]string) (*models.AccessGroup, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", fmt.Sprintf("/accessgroup/%d", id)).
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	return convertToStruct[models.AccessGroup](resp)
}
