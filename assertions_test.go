package assertions_test

import (
	"testing"

	"github.com/achedges/go-assertions"
)

func Test_True(t *testing.T) {
	mockTest := testing.T{}

	assertions.True(true, &mockTest)
	if mockTest.Failed() {
		t.Errorf("Expected passing but failed")
	}

	assertions.True(false, &mockTest)
	if !mockTest.Failed() {
		t.Errorf("Expected failing but passed")
	}
}

func Test_False(t *testing.T) {
	mockTest := testing.T{}

	assertions.False(false, &mockTest)
	if mockTest.Failed() {
		t.Errorf("Expected passing but failed")
	}

	assertions.False(true, &mockTest)
	if !mockTest.Failed() {
		t.Errorf("Expected failing but passed")
	}
}

func Test_EqualInts(t *testing.T) {
	mockTest := testing.T{}

	assertions.EqualInts(12345, 12345, &mockTest)
	assertions.EqualInts(-12345, -12345, &mockTest)
	if mockTest.Failed() {
		t.Errorf("Expected passing but failed")
	}

	assertions.EqualInts(12345, 12344, &mockTest)
	if !mockTest.Failed() {
		t.Errorf("Expected failing but passed")
	}
}

func Test_EqualUints(t *testing.T) {
	mockTest := testing.T{}
	a := uint32(12345)
	b := uint32(12345)
	c := uint32(12344)

	assertions.EqualUints(a, b, &mockTest)
	if mockTest.Failed() {
		t.Errorf("Expected passing but failed")
	}

	assertions.EqualUints(a, c, &mockTest)
	if !mockTest.Failed() {
		t.Errorf("Expected failing but passed")
	}
}

func Test_EqualFloats(t *testing.T) {
	mockTest := testing.T{}
	a := float32(1.2345)
	b := float32(1.2345)
	c := float32(1.2344)

	assertions.EqualFloats(a, b, &mockTest)
	if mockTest.Failed() {
		t.Errorf("Expected passing but failed")
	}

	assertions.EqualFloats(a, c, &mockTest)
	if !mockTest.Failed() {
		t.Errorf("Expected failing but passed")
	}
}

func Test_CloseEnough(t *testing.T) {
	mockTest := testing.T{}

	assertions.CloseEnough(12.34, 12.343, 0.01, &mockTest)
	assertions.CloseEnough(12.345, 12.346, 0.001, &mockTest)
	assertions.CloseEnough(12.3456, 12.34567, 0.0001, &mockTest)
	if mockTest.Failed() {
		t.Errorf("Expected passing but failed")
	}

	assertions.CloseEnough(12.34, 13.12, 0.01, &mockTest)
	if !mockTest.Failed() {
		t.Errorf("Expected failed but passing")
	}
}

func Test_EqualStrings(t *testing.T) {
	mockTest := testing.T{}

	assertions.EqualStrings("asdf", "asdf", &mockTest)
	if mockTest.Failed() {
		t.Errorf("Expected passing but faile")
	}

	assertions.EqualStrings("asdf", "qwer", &mockTest)
	if !mockTest.Failed() {
		t.Errorf("Expected failing but passed")
	}
}
