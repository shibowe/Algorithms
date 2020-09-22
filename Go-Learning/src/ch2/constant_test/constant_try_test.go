package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday)
}
