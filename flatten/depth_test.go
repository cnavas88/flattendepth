package flatten

import (
	"testing"
)

func TestFirstDepth(t *testing.T) {
	nestedArray := []interface{}{
		[]interface{}{
			10,
			20,
			30,
		},
		40,
	}

	expectedDepth := 2

	depth, err := Depth(nestedArray)

	if depth != expectedDepth {
		t.Errorf("Expected depth %v, but got %v", expectedDepth, depth)
	}

	if err != nil {
		t.Errorf("Expected not error, but got %v", err)
	}
}

func TestSecondDepth(t *testing.T) {
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

	expectedDepth := 3

	depth, err := Depth(nestedArray)

	if expectedDepth != depth {
		t.Errorf("Expected depth %v, but got %v", expectedDepth, depth)
	}

	if err != nil {
		t.Errorf("Expected not error, but got %v", err)
	}
}

func TestThirdDepthArray(t *testing.T) {
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

	expectedDepth := 5

	depth, err := Depth(nestedArray)

	if expectedDepth != depth {
		t.Errorf("Expected array %v, but got %v", expectedDepth, depth)
	}

	if err != nil {
		t.Errorf("Expected not error, but got %v", err)
	}
}

func TestFourthDepthArray(t *testing.T) {
	nestedArray := []interface{}{
		10,
		20,
		30,
	}

	expectedDepth := 1

	depth, err := Depth(nestedArray)

	if expectedDepth != depth {
		t.Errorf("Expected depth %v, but got %v", expectedDepth, depth)
	}

	if err != nil {
		t.Errorf("Expected not error, but got %v", err)
	}
}

func TestErrorArrayEmptyToDepth(t *testing.T) {
	_, err := Depth(nil)

	if err.Error() != "Invalid array" {
		t.Errorf("Expected error: %v", err.Error())
	}
}

func TestErrorNotArrayOrSliceToDepth(t *testing.T) {
	_, err := Depth("45")

	if err.Error() != "Not supported value type" {
		t.Errorf("Expected error: %v", err.Error())
	}
}
