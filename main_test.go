package main

import (
	"testing"
)

func TestHi(t *testing.T) {
	if Hi() != "Hi" {
		t.Errorf("This is not Hi")
	}
}
