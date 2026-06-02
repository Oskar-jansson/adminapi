package services

import (
	"bufio"
	"context"
	"encoding/json"
	"net/http"

	"github.com/Oskar-jansson/adminapi/models"
)

type EventService struct {
	sc clientInterface
}

// start eventstream.
// holds until error.
// context is honored
func (c *EventService) Stream(ctx context.Context, onEvent func(event *models.Event), params map[string]string) error {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", "/eventstream").
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return err
	}

	return reader[*models.Event](ctx, resp, onEvent)
}

// helper function to read incomming stream
func reader[T any](ctx context.Context, resp *http.Response, onEvent func(event T)) error {
	defer func() {
		resp.Body.Close()
	}()

	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		var event T
		if err := json.Unmarshal([]byte(line), &event); err != nil {
			continue
		}

		onEvent(event)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil

}

func (c *EventService) Request(ctx context.Context, params map[string]string) error {
	resp, err := NewRequest(ctx, c.sc).
		Method("GET", "/accessgroup").
		Query(params).
		WithBearer().
		Do()

	if err != nil {
		return err
	}
	_, err = convertToStruct[models.AccessGroupList](resp)
	return err
}
