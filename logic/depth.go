package logic

// Depth return depth of array or slice
func Depth(array interface{}) (int, error) {
	if err := controlErrors(array); err != nil {
		return -1, err
	}
	return doDepth(array, 0), nil
}

func doDepth(array interface{}, depth int) int {
	var subDepth int

	s := array.([]interface{})

	maxDepth := depth

	for _, subArray := range s {
		switch v := subArray.(type) {
		case []interface{}:
			subDepth = doDepth(v, depth+1)
			if subDepth > maxDepth {
				maxDepth = subDepth
			}

		}
	}
	return depth
}