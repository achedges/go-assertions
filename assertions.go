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
	if !condition {
		t.Error()
	}
}

func False(condition bool, t *testing.T) {
	if condition {
		t.Error()
	}
}

func EqualInts[T Int](expected T, actual T, t *testing.T) {
	if expected != actual {
		t.Error()
	}
}

func EqualUints[T UInt](expected T, actual T, t *testing.T) {
	if expected != actual {
		t.Error()
	}
}

func EqualFloats[T Float](expected T, actual T, t *testing.T) {
	if expected != actual {
		t.Error()
	}
}

func CloseEnough[T Float](expected T, actual T, delta T, t *testing.T) {
	if (math.Abs(float64(expected) - float64(actual))) > float64(delta) {
		t.Error()
	}
}
