package eziql

import (
	"testing"
)

func TestNoteMarshalling(t *testing.T) {
	a := "hello"
	if a != "hello" {
		t.Error("problem")
	}
}
