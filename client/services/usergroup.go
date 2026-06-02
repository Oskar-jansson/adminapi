package services

import (
	"context"
	"fmt"

	"github.com/Oskar-jansson/adminapi/models"
)

type UserGroupService struct {
	sc clientInterface
}

func (c *UserGroupService) Get(ctx context.Context, id int, params map[string]string) (*models.UserGroup, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", fmt.Sprintf("/usergroup/%d", id)).
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.UserGroup](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *UserGroupService) List(ctx context.Context, params map[string]string) (*models.UserGroupList, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", "/usergroup").
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.UserGroupList](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *UserGroupService) Create(ctx context.Context, newUserGroup models.DepartmentInput) (*models.UserGroup, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("POST", "/usergroup").
		Body(newUserGroup).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.UserGroup](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *UserGroupService) Edit(ctx context.Context, id int, changes models.DepartmentInput) (*models.UserGroup, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("PATCH", fmt.Sprintf("/usergroup/%d", id)).
		Body(changes).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.UserGroup](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *UserGroupService) Delete(ctx context.Context, id int) error {
	resp, err := NewRequest(ctx, c.sc).
		Method("DELETE", fmt.Sprintf("/usergroup/%d", id)).
		WithBearer().
		Do()

	if err != nil {
		return err
	}

	_, err = convertToStruct[struct{}](resp)

	return err
}
