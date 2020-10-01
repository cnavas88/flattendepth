package logic

// Flatten enter point to algorithm.
func Flatten(array interface{}) ([]int, error) {
	if err := controlErrors(array); err != nil {
		return nil, err
	}
	return doFlatten(array, []int{}), nil
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
