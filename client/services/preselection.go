package services

import (
	"context"
	"fmt"

	"github.com/Oskar-jansson/adminapi/models"
)

type PreselectionService struct {
	sc clientInterface
}

func (c *PreselectionService) Get(ctx context.Context, id int, params map[string]string) (*models.Preselection, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", fmt.Sprintf("/preselection/%d", id)).
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.Preselection](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *PreselectionService) List(ctx context.Context, params map[string]string) (*models.PreselectionList, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", "/preselection").
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.PreselectionList](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
