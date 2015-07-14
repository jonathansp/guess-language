package utils

import "testing"

func TestInFunction(t *testing.T) {

	b := []string{"a", "b", "c"}
	if has := In("a", b); !has {
		t.Fatalf("Expected true, got false.")
	}
	if has := In("z", b); has {
		t.Fatalf("Expected false, got true.")
	}
}

func TestHasOneFunction(t *testing.T) {
	a := []string{"x", "y"}
	b := []string{"a", "b", "c"}
	if has := HasOne(a, b); has {
		t.Fatalf("Expected false, got true.")
	}

}
