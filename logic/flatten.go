package logic

import (
	"errors"
	"fmt"
)

// Flatten enter point to algorithm.
func Flatten(array interface{}) ([]int, error) {
	if array == nil {
		return nil, errors.New("Invalid array")
	}

	switch v := array.(type) {
	case []int, []interface{}:
		return doFlatten(array, []int{}), nil

	default:
		return nil, fmt.Errorf("Not supported value type %T", v)
	}
}

func doFlatten(array interface{}, acc []int) []int {
	s := array.([]interface{})

	for _, subArray := range s {
		switch v := subArray.(type) {
		case []interface{}:
			acc = doFlatten(v, acc)

		case int:
			acc = append(acc, v)
		}
	}

	return acc
}

// package logic

// import (
// 	"errors"
// 	"fmt"
// )

// func Depth(array interface{}) (int, error) {
// 	if array == nil {
// 		return nil, errors.New("Invalid array")
// 	}

// 	switch v := array.(type) {
// 	case []int, []interface{}:
// 		return deep(array, 1), nil

// 	default:
// 		return nil, fmt.Errorf("Not supported value type %T", v)
// 	}
// }

// func doDepth(array []interface{}, depth int) int {
// 	var subDepth int

// 	s := array.([]interface{})

// 	maxDepth := depth

// 	for _, subArray := range s {
// 		switch v := subArray.(type) {
// 		case []interface{}:
// 			subDepth = doDepth(v, depth + 1)
// 			if subDepth > maxDepth {
// 				maxDepth = subDepth
// 			}

// 	}

// 	return depth
// }
