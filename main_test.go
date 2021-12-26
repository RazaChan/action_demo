package main

import (
	"testing"
)

func TestCat(t *testing.T) {
	saying := Cat()
	if saying != "Meow~~~~" {
		t.Errorf("Cat say %s", saying)
	}
}
