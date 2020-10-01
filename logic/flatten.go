package logic

// Flatten enter point to algorithm.
func Flatten(array interface{}) ([]int, error) {
	if err := isValidArray(array); err != nil {
		return nil, err
	}
	return doFlatten(array, []int{}), nil
}

func doFlatten(array interface{}, acc []int) []int {
	s := array.([]interface{})

	for _, subArray := range s {
		switch v := subArray.(type) {
		case int:
			acc = append(acc, v)

		case []int:
			acc = append(acc, v...)

		case []interface{}:
			acc = doFlatten(v, acc)
		}
	}

	return acc
}
