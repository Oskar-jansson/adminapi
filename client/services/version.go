package services

import (
	"context"

	"github.com/Oskar-jansson/adminapi/models"
)

type VersionService struct {
	sc clientInterface
}

func NewVersionService(c clientInterface) *VersionService {
	return &VersionService{sc: c}
}

func (c *VersionService) Get(ctx context.Context) (*models.Version, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", "version").
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.Version](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil

}
