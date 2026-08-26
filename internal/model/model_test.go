package model

import "testing"

func TestModelsValid(t *testing.T) {
	if !NewRecord("r", "f", "s", "view").Valid() {
		t.Fatal()
	}
	if !NewProfile("p", "n").Valid() {
		t.Fatal()
	}
	if !NewEvent("e", "r", "view", "s").Valid() {
		t.Fatal()
	}
	if !NewAudit("a", "Record", "r", "view", "s").Valid() {
		t.Fatal()
	}
}
