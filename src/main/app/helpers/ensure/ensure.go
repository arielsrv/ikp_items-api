package ensure

import (
	"ikp_items-api/src/main/app/server/errors"
	"net/http"
)

func NotEmpty(value string, message string) error {
	if value == "" {
		return errors.NewError(http.StatusBadRequest, message)
	}
	return nil
}

func Int(value int, message string) error {
	if value < 1 {
		return errors.NewError(http.StatusBadRequest, message)
	}
	return nil
}

func Int64(value int64, message string) error {
	if value < 1 {
		return errors.NewError(http.StatusBadRequest, message)
	}

	return nil
}

type SafeEnum interface {
	IsValid() bool
}

func Enum(safeEnum SafeEnum, message string) error {
	if !safeEnum.IsValid() {
		return errors.NewError(http.StatusBadRequest, message)
	}

	return nil
}
