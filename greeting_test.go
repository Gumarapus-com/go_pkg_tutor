package greeting

import "testing"

func TestGreet(t *testing.T) {
	if Greet("My name") != nil {
		t.Error("Expecting nil")
	}
}
