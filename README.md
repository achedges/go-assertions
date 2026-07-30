# Assertions
A simple unit test assertions library for Go

This library does nothing more than add syntax-sugar for working with standard Go unit tests.

Instead of writing this:
```go
func TestThings(t *testing.T) {
	var a int = 10
	var b int = 10
	if a != b {
		t.Error()
	}
}
```

...you can write this:
```go
func TestThings(t *testing.T) {
	var a int = 10
	var b int = 10
	assertions.EqualInts(a, b, t)
}
```