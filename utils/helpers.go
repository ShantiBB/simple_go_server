package utils

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func ReadIDParam(r *http.Request) (int64, error) {
	param := r.PathValue("id")

	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil || id < 1 {
		errInfo := fmt.Sprintf("Parameter isn't valid: %s", param)
		return 0, errors.New(errInfo)
	}

	return id, nil
}
