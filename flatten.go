package main

import (
	"errors"
)

// Flatten enter point to algorithm.
func Flatten(array interface{}) ([]int, error) {
	if err := isValidArray(array); err != nil {
		return nil, err
	}

	result, err := doFlatten(array, []int{})

	if err != nil {
		return nil, err
	}
	return result, nil
}

func doFlatten(array interface{}, acc []int) ([]int, error) {
	var err error
	s := array.([]interface{})

	for _, subArray := range s {
		switch v := subArray.(type) {
		case int:
			acc = append(acc, v)

		case []int:
			acc = append(acc, v...)

		case []interface{}:
			acc, err = doFlatten(v, acc)
			if err != nil {
				return nil, err
			}

		default:
			return nil, errors.New("Value not supported")
		}
	}

	return acc, nil
}
