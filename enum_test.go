package enum_test

import (
	"testing"

	"github.com/ElioenaiFerrari/enum"
)

func TestIsValid(t *testing.T) {
	values := []string{
		"Gandalf",
		"Frodo",
		"Sauron",
	}

	e := enum.New(values...)

	for _, v := range values {
		if e.IsValid(v) == false {
			t.Error("expected true, got false")
		}
	}

	if e.IsValid("Aragorm") == true {
		t.Error("expected false, got true")
	}
}
