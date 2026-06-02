package services

import (
	"context"
	"fmt"

	"github.com/Oskar-jansson/adminapi/models"
)

type UserService struct {
	sc clientInterface
}

func (c *UserService) Get(ctx context.Context, id int, params map[string]string) (*models.User, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", fmt.Sprintf("/user/%d", id)).
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.User](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *UserService) List(ctx context.Context, params map[string]string) (*models.UserList, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", "/user").
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.UserList](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *UserService) Create(ctx context.Context, newUser models.UserInput) (*models.User, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("POST", "/user").
		Body(newUser).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.User](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *UserService) Edit(ctx context.Context, id int, changes models.UserInput) (*models.User, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("PATCH", fmt.Sprintf("/user/%d", id)).
		Body(changes).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.User](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *UserService) Delete(ctx context.Context, id int) error {
	resp, err := NewRequest(ctx, c.sc).
		Method("DELETE", fmt.Sprintf("/user/%d", id)).
		WithBearer().
		Do()

	if err != nil {
		return err
	}

	_, err = convertToStruct[struct{}](resp)

	return err
}

func (c *UserService) AssignAccessGroup(ctx context.Context, userId int, accessGroupId int) error {
	resp, err := NewRequest(ctx, c.sc).
		Method("PUT", fmt.Sprintf("/user/%d/accessgroup/%d", userId, accessGroupId)).
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

func (c *UserService) RemoveAccessGroup(ctx context.Context, userId int, accessGroupId int) error {
	resp, err := NewRequest(ctx, c.sc).
		Method("DELETE", fmt.Sprintf("/user/%d/accessgroup/%d", userId, accessGroupId)).
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
