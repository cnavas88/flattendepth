package logic

import (
	"errors"
	"log"
	"reflect"
)

// Flatten enter point to algorithm, i am using []float64 return type because Unmarshal function
// decode the json numbers to float64 type.
func Flatten(array interface{}) ([]float64, int8, error) {
	if array == nil {
		return nil, -1, errors.New("The structure to flatten is empty")
	}

	t := reflect.TypeOf(array)

	if t.Kind() != reflect.Slice && t.Kind() != reflect.Array {
		return nil, -1, errors.New("The structure to flatten is not and array or slice")
	}

	acc, depth := doFlatten(array, []float64{}, 1)

	return acc, depth, nil
}

func doFlatten(array interface{}, acc []float64, depth int8) ([]float64, int8) {
	var subDepth int8

	s := array.([]interface{})

	maxDepth := depth

	for _, subArray := range s {
		switch v := subArray.(type) {
		case []interface{}:
			acc, subDepth = doFlatten(v, acc, depth+1)
			if subDepth > maxDepth {
				maxDepth = subDepth
			}

		case float64:
			acc = append(acc, v)

		default:
			log.Printf("The data %v not flatten. Type invalid.", v)
		}
	}

	return acc, maxDepth
}
