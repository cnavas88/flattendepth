package logic

import (
	"reflect"
	"testing"
)

func TestFirstFlattenArray(t *testing.T) {
	nestedArray := []interface{}{
		[]interface{}{
			10.0,
			20.0,
			30.0,
		},
		40.0,
	}

	expectedArray := []float64{10, 20, 30, 40}

	arrayFlatten, deep, err := Flatten(nestedArray)

	if reflect.DeepEqual(arrayFlatten, expectedArray) == false {
		t.Errorf("Expected array %v, but got %v", expectedArray, arrayFlatten)
	}

	if deep != 2 {
		t.Errorf("Expected deep 2, but got %v", deep)
	}

	if err != nil {
		t.Errorf("Expected not error, but got %v", err)
	}
}

func TestSecondFlattenArray(t *testing.T) {
	nestedArray := []interface{}{
		[]interface{}{
			10.0,
			20.0,
			[]interface{}{
				30.0,
			},
		},
		40.0,
	}

	expectedArray := []float64{10, 20, 30, 40}

	arrayFlatten, deep, err := Flatten(nestedArray)

	if reflect.DeepEqual(arrayFlatten, expectedArray) == false {
		t.Errorf("Expected array %v, but got %v", expectedArray, arrayFlatten)
	}

	if deep != 3 {
		t.Errorf("Expected deep 3, but got %v", deep)
	}

	if err != nil {
		t.Errorf("Expected not error, but got %v", err)
	}
}

func TestThirdFlattenArray(t *testing.T) {
	nestedArray := []interface{}{
		[]interface{}{
			10.0,
			[]interface{}{
				[]interface{}{
					20.0,
					[]interface{}{30.0},
				},
			},
			[]interface{}{40.0},
		},
	}

	expectedArray := []float64{10, 20, 30, 40}

	arrayFlatten, deep, err := Flatten(nestedArray)

	if reflect.DeepEqual(arrayFlatten, expectedArray) == false {
		t.Errorf("Expected array %v, but got %v", expectedArray, arrayFlatten)
	}

	if deep != 5 {
		t.Errorf("Expected deep 5, but got %v", deep)
	}

	if err != nil {
		t.Errorf("Expected not error, but got %v", err)
	}
}

func TestFourthFlattenArray(t *testing.T) {
	nestedArray := []interface{}{
		10.0,
		20.0,
		30.0,
	}

	expectedArray := []float64{10, 20, 30}

	arrayFlatten, deep, err := Flatten(nestedArray)

	if reflect.DeepEqual(arrayFlatten, expectedArray) == false {
		t.Errorf("Expected array %v, but got %v", expectedArray, arrayFlatten)
	}

	if deep != 1 {
		t.Errorf("Expected deep 1, but got %v", deep)
	}

	if err != nil {
		t.Errorf("Expected not error, but got %v", err)
	}
}

func TestErrorArrayEmpty(t *testing.T) {
	_, _, err := Flatten(nil)

	if err.Error() != "The structure to flatten is empty" {
		t.Errorf("Expected error: The structure to flatten is empty")
	}
}

func TestErrorNotArrayOrSlice(t *testing.T) {
	_, _, err := Flatten("45")

	if err.Error() != "The structure to flatten is not and array or slice" {
		t.Errorf("Expected error: The structure to flatten is not and array or slice")
	}
}
