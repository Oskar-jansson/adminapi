package services

import (
	"context"
	"fmt"

	"github.com/Oskar-jansson/adminapi/models"
)

type FunctionService struct {
	sc clientInterface
}

func (c *FunctionService) Get(ctx context.Context, id int, params map[string]string) (*models.Function, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", fmt.Sprintf("/function/%d", id)).
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.Function](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *FunctionService) List(ctx context.Context, params map[string]string) (*models.FunctionList, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", "/function").
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.FunctionList](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *FunctionService) Create(ctx context.Context, newFunction models.DepartmentInput) (*models.Function, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("POST", "/function").
		Body(newFunction).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.Function](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *FunctionService) Edit(ctx context.Context, id int, changes models.DepartmentInput) (*models.Function, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("PATCH", fmt.Sprintf("/function/%d", id)).
		Body(changes).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.Function](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *FunctionService) Delete(ctx context.Context, id int) error {
	resp, err := NewRequest(ctx, c.sc).
		Method("DELETE", fmt.Sprintf("/function/%d", id)).
		WithBearer().
		Do()

	if err != nil {
		return err
	}

	_, err = convertToStruct[struct{}](resp)

	return err
}
