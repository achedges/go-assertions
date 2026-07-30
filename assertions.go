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

func EqualInts[T Int](a T, b T, t *testing.T) {
	if a != b {
		t.Error()
	}
}

func EqualUints[T UInt](a T, b T, t *testing.T) {
	if a != b {
		t.Error()
	}
}

func EqualFloats[T Float](a T, b T, t *testing.T) {
	if a != b {
		t.Error()
	}
}

func CloseEnough[T Float](a T, b T, delta T, t *testing.T) {
	if (math.Abs(float64(a) - float64(b))) > float64(delta) {
		t.Error()
	}
}
