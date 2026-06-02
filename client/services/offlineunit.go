package services

import (
	"context"
	"fmt"

	"github.com/Oskar-jansson/adminapi/models"
)

type OfflineUnitService struct {
	sc clientInterface
}

func (c *OfflineUnitService) Get(ctx context.Context, id int, params map[string]string) (*models.OfflineUnit, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", fmt.Sprintf("/offlineunit/%d", id)).
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.OfflineUnit](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *OfflineUnitService) List(ctx context.Context, params map[string]string) (*models.OfflineUnitList, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", "/offlineunit").
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.OfflineUnitList](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *OfflineUnitService) StepAccessVersion(ctx context.Context, id int) error {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", fmt.Sprintf("/offlineunit/%d/stepaccessversion", id)).
		WithBearer().
		Do()

	if err != nil {
		return err
	}

	_, err = convertToStruct[struct{}](resp)

	return err
}
