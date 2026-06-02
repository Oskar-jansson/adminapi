package services

import (
	"context"
	"fmt"

	"github.com/Oskar-jansson/adminapi/models"
)

type MachineGroupService struct {
	sc clientInterface
}

func (c *MachineGroupService) Get(ctx context.Context, id int, params map[string]string) (*models.MachineGroup, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", fmt.Sprintf("/machinegroup/%d", id)).
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.MachineGroup](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *MachineGroupService) List(ctx context.Context, params map[string]string) (*models.MachineGroupList, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", "/machinegroup").
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.MachineGroupList](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
