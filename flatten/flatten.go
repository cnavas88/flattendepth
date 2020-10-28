package flatten

// Flatten enter point to algorithm.
func Flatten(array interface{}) ([]int, error) {
	var flatten []int

	if err := isValidArray(array); err != nil {
		return nil, err
	}

	if err := doFlatten(array, &flatten); err != nil {
		return nil, err
	}

	return flatten, nil
}

func doFlatten(array interface{}, acc *[]int) error {
	s := array.([]interface{})

	for _, subArray := range s {
		switch v := subArray.(type) {
		case int:
			*acc = append(*acc, v)

		case []int:
			*acc = append(*acc, v...)

		case []interface{}:
			if err := doFlatten(v, acc); err != nil {
				return err
			}

		default:
			return ErrNotSupportType
		}
	}

	return nil
}
