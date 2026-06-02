package services

import (
	"context"

	"github.com/Oskar-jansson/adminapi/models"
)

type SettingService struct {
	sc clientInterface
}

func (c *SettingService) List(ctx context.Context, params map[string]string) (*models.SettingList, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", "/setting").
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.SettingList](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
