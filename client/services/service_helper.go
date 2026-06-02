package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync/atomic"

	"github.com/Oskar-jansson/adminapi/models"
)

var (
	AllowUnknownFields atomic.Bool
	AllowTrailingData  atomic.Bool
)

func init() {
	AllowUnknownFields.Store(true)
	AllowTrailingData.Store(true)
}

func convertToStruct[T any](resp *http.Response) (*T, error) {

	if resp == nil {
		return nil, ErrClientNilResponse
	}

	defer func() {
		_ = resp.Body.Close() // disregard error on reader close
	}()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if isErrorStatus(resp.StatusCode) {

		// attempt to parse error as Api-ErrorMessage
		var errorMessage models.ErrorMessage
		err := json.Unmarshal(raw, &errorMessage)
		if err == nil {
			return nil, parseApiErrorCode(&errorMessage)
		}

		// attempt to return response status as error
		if err != nil && resp.Status != "" {
			return nil, fmt.Errorf("%s", resp.Status)
		}

		// fallback to unkownError
		return nil, ErrClientUnkownError

	}

	var obj *T

	// If the body is empty, return a zero value instead of unmarshalling
	if len(bytes.TrimSpace(raw)) == 0 {
		return obj, nil
	}

	err = json.Unmarshal(raw, &obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func decode[T any](raw []byte) (T, error) {
	var obj T

	dec := json.NewDecoder(bytes.NewReader(raw))

	if !AllowUnknownFields.Load() {
		dec.DisallowUnknownFields()
	}

	if err := dec.Decode(&obj); err != nil {
		return obj, errors.Join(err, ErrClientDecodeError)
	}

	if !AllowTrailingData.Load() {
		var extra any
		if err := dec.Decode(&extra); err != io.EOF {
			return obj, ErrClientTrailingData
		}
	}

	return obj, nil
}

// Helper to centralize logic for non-successful HTTP status codes.
func isErrorStatus(v int) bool {
	return v < 200 || v >= 300
}
