package testing

import "testing"

// AssertEquals calls t.Fatalf if want != got.
func AssertEquals(t *testing.T, want, got interface{}) {
	if want != got {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}

// AssertNotEquals calls t.Fatalf if want == got.
func AssertNotEquals(t *testing.T, doNotWant, got interface{}) {
	if doNotWant == got {
		t.Fatalf("Did not want %v, but got %v", doNotWant, got)
	}
}
