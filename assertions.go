package assertions

import (
	"math"
	"testing"
)

type Int interface {
	int | int8 | int16 | int32 | int64
}

type UInt interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type Float interface {
	float32 | float64
}

func True(condition bool, t *testing.T) {
	t.Helper()
	if !condition {
		t.Errorf("Expected %t, was %t", true, condition)
	}
}

func False(condition bool, t *testing.T) {
	t.Helper()
	if condition {
		t.Errorf("Expected %t, was %t", false, condition)
	}
}

func EqualInts[T Int](expected T, actual T, t *testing.T) {
	t.Helper()
	if expected != actual {
		t.Errorf("Expected %d, was %d", expected, actual)
	}
}

func EqualUints[T UInt](expected T, actual T, t *testing.T) {
	t.Helper()
	if expected != actual {
		t.Errorf("Expected %d, was %d", expected, actual)
	}
}

func EqualFloats[T Float](expected T, actual T, t *testing.T) {
	t.Helper()
	if expected != actual {
		t.Errorf("Expected %.6f, was %.6f", expected, actual)
	}
}

func CloseEnough[T Float](expected T, actual T, delta T, t *testing.T) {
	t.Helper()
	if (math.Abs(float64(expected) - float64(actual))) > float64(delta) {
		t.Errorf("Expected %.6f (+/- %.6f), was %6.f", expected, delta, actual)
	}
}

func EqualStrings(expected string, actual string, t *testing.T) {
	t.Helper()
	if expected != actual {
		t.Errorf("Expected [%s], was [%s]", expected, actual)
	}
}
