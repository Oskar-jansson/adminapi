package services

import (
	"context"
	"fmt"

	"github.com/Oskar-jansson/adminapi/models"
)

type CardService struct {
	sc clientInterface
}

func (c *CardService) List(ctx context.Context, params map[string]string) (*models.CardList, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", "/card").
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.CardList](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *CardService) Get(ctx context.Context, id int, params map[string]string) (*models.Card, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", fmt.Sprintf("/card/%d", id)).
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.Card](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *CardService) Create(ctx context.Context, newCard models.CardInput) (*models.Card, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("POST", "/card").
		Body(newCard).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.Card](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *CardService) Edit(ctx context.Context, id int, changes models.CardInput) (*models.Card, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("PATCH", fmt.Sprintf("/card/%d", id)).
		Body(changes).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.Card](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *CardService) Delete(ctx context.Context, id int, params map[string]string) error {
	resp, err := NewRequest(ctx, c.sc).
		Method("DELETE", fmt.Sprintf("/card/%d", id)).
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return err
	}

	_, err = convertToStruct[models.Card](resp)

	return err
}

//
//	Accessgroup for card
//

func (c *CardService) AssignAccessGroup(ctx context.Context, cardId int, accessGroupId int) error {
	resp, err := NewRequest(ctx, c.sc).
		Method("PUT", fmt.Sprintf("/card/%d/accessgroup/%d", cardId, accessGroupId)).
		WithBearer().
		Do()

	if err != nil {
		return err
	}

	_, err = convertToStruct[struct{}](resp)
	if err != nil {
		return err
	}

	return nil
}

func (c *CardService) RemoveAccessGroup(ctx context.Context, cardId int, accessGroupId int) error {
	resp, err := NewRequest(ctx, c.sc).
		Method("DELETE", fmt.Sprintf("/card/%d/accessgroup/%d", cardId, accessGroupId)).
		WithBearer().
		Do()

	if err != nil {
		return err
	}

	_, err = convertToStruct[struct{}](resp)
	if err != nil {
		return err
	}

	return nil
}
