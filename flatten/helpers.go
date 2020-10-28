package flatten

import (
	"errors"
)

// ErrInvalidArray array can't be nil
var ErrInvalidArray = errors.New("Invalid array")

// ErrNotSupportType define error to check supported values
var ErrNotSupportType = errors.New("Not supported value type")

func isValidArray(array interface{}) error {
	if array == nil {
		return ErrInvalidArray
	}

	switch array.(type) {
	case []int, []interface{}:
		return nil

	default:
		return ErrNotSupportType
	}
}
