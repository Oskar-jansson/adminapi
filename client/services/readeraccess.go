package services

import (
	"context"
	"fmt"

	"github.com/Oskar-jansson/adminapi/models"
)

type ReaderAccessService struct {
	sc clientInterface
}

func (c *ReaderAccessService) List(ctx context.Context, params map[string]string) (*models.ReaderAccessList, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", "/readeraccess").
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.ReaderAccessList](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *ReaderAccessService) Create(ctx context.Context, cardId int, newReaderAccess models.ReaderAccessInput) (*models.ReaderAccess, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("POST", fmt.Sprintf("/card/%d/readeraccess", cardId)).
		Body(newReaderAccess).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.ReaderAccess](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *ReaderAccessService) Delete(ctx context.Context, cardId, unitId int) error {
	resp, err := NewRequest(ctx, c.sc).
		Method("DELETE", fmt.Sprintf("/card/%d/readeraccess/%d", cardId, unitId)).
		WithBearer().
		Do()

	if err != nil {
		return err
	}

	_, err = convertToStruct[struct{}](resp)

	return err
}
