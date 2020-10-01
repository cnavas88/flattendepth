package main

import (
	"errors"
	"fmt"
)

func isValidArray(array interface{}) error {
	if array == nil {
		return errors.New("Invalid array")
	}

	switch v := array.(type) {
	case []int, []interface{}:
		return nil

	default:
		return fmt.Errorf("Not supported value type %T", v)
	}
}
