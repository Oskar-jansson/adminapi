package services

import (
	"context"
	"fmt"

	"github.com/Oskar-jansson/adminapi/models"
)

type DepartmentService struct {
	sc clientInterface
}

func (c *DepartmentService) Get(ctx context.Context, id int, params map[string]string) (*models.Department, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", fmt.Sprintf("/department/%d", id)).
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.Department](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *DepartmentService) List(ctx context.Context, params map[string]string) (*models.DepartmentList, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", "/department").
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.DepartmentList](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *DepartmentService) Create(ctx context.Context, newDepartment models.DepartmentInput) (*models.Department, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("POST", "/department").
		Body(newDepartment).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.Department](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *DepartmentService) Edit(ctx context.Context, id int, changes models.DepartmentInput) (*models.Department, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("PATCH", fmt.Sprintf("/department/%d", id)).
		Body(changes).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.Department](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *DepartmentService) Delete(ctx context.Context, id int) error {
	resp, err := NewRequest(ctx, c.sc).
		Method("DELETE", fmt.Sprintf("/department/%d", id)).
		WithBearer().
		Do()

	if err != nil {
		return err
	}

	_, err = convertToStruct[struct{}](resp)

	return err
}
