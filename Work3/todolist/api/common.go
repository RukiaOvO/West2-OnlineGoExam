package api

import (
	"encoding/json"

	"todolist/pkg/ctl"
	"todolist/pkg/e"
)

func ErrorResponse(err error) *ctl.Response {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return ctl.RespError(err, e.ErrorUnmarshalJson)
	}
	return ctl.RespError(err, e.Error)
}
