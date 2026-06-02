package services

import (
	"context"
	"fmt"

	"github.com/Oskar-jansson/adminapi/models"
)

type MachineGroupTypeService struct {
	sc clientInterface
}

func (c *MachineGroupTypeService) Get(ctx context.Context, id int, params map[string]string) (*models.MachineGroupType, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", fmt.Sprintf("/machinegrouptype/%d", id)).
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.MachineGroupType](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (c *MachineGroupTypeService) List(ctx context.Context, params map[string]string) (*models.MachineGroupTypeList, error) {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", "/machinegrouptype").
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return nil, err
	}

	obj, err := convertToStruct[models.MachineGroupTypeList](resp)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
