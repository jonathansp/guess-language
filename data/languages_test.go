package data

import "testing"

func TestSizeMap(t *testing.T) {
	total := 94
	if size := len(Languages); size != total {
		t.Fatalf("Expected %d, got %d.", size, total)
	}
}
