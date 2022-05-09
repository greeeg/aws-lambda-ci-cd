package message

import (
	"testing"
)

func TestGreet(t *testing.T) {
	actualString := Greet("Gopher")
	expectedString := "Hello Gopher"

	if actualString != expectedString {
		t.Errorf("Expected String(%s) is not same as"+
			" actual string (%s)", expectedString, actualString)
	}
}
