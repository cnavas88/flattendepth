package flatten

import (
	"reflect"
	"testing"
)

func TestFirstFlattenArray(t *testing.T) {
	nestedArray := []interface{}{
		[]interface{}{
			10,
			20,
			30,
		},
		40,
	}

	expectedArray := []int{10, 20, 30, 40}

	arrayFlatten, err := Flatten(nestedArray)

	if reflect.DeepEqual(arrayFlatten, expectedArray) == false {
		t.Errorf("Expected array %v, but got %v", expectedArray, arrayFlatten)
	}

	if err != nil {
		t.Errorf("Expected not error, but got %v", err)
	}
}

func TestSecondFlattenArray(t *testing.T) {
	nestedArray := []interface{}{
		[]interface{}{
			10,
			20,
			[]interface{}{
				30,
			},
		},
		40,
	}

	expectedArray := []int{10, 20, 30, 40}

	arrayFlatten, err := Flatten(nestedArray)

	if reflect.DeepEqual(arrayFlatten, expectedArray) == false {
		t.Errorf("Expected array %v, but got %v", expectedArray, arrayFlatten)
	}

	if err != nil {
		t.Errorf("Expected not error, but got %v", err)
	}
}

func TestThirdFlattenArray(t *testing.T) {
	nestedArray := []interface{}{
		[]interface{}{
			10,
			[]interface{}{
				[]interface{}{
					20,
					[]interface{}{30},
				},
			},
			[]int{40, 50, 60, 70},
		},
	}

	expectedArray := []int{10, 20, 30, 40, 50, 60, 70}

	arrayFlatten, err := Flatten(nestedArray)

	if reflect.DeepEqual(arrayFlatten, expectedArray) == false {
		t.Errorf("Expected array %v, but got %v", expectedArray, arrayFlatten)
	}

	if err != nil {
		t.Errorf("Expected not error, but got %v", err)
	}
}

func TestFourthFlattenArray(t *testing.T) {
	nestedArray := []interface{}{
		10,
		20,
		30,
	}

	expectedArray := []int{10, 20, 30}

	arrayFlatten, err := Flatten(nestedArray)

	if reflect.DeepEqual(arrayFlatten, expectedArray) == false {
		t.Errorf("Expected array %v, but got %v", expectedArray, arrayFlatten)
	}

	if err != nil {
		t.Errorf("Expected not error, but got %v", err)
	}
}

func TestErrorArrayEmpty(t *testing.T) {
	_, err := Flatten(nil)

	if err.Error() != "Invalid array" {
		t.Errorf("Expected error: %v", err.Error())
	}
}

func TestErrorNotArrayOrSlice(t *testing.T) {
	_, err := Flatten("45")

	if err.Error() != "Not supported value type" {
		t.Errorf("Expected error: %v", err.Error())
	}
}

func TestErrorNotValidValueInsideArray(t *testing.T) {
	nestedArray := []interface{}{
		10,
		20,
		"30",
	}

	_, err := Flatten(nestedArray)

	if err.Error() != "Not supported value type" {
		t.Errorf("Expected error: %v", err.Error())
	}
}
