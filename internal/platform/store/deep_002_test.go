package store

import "testing"

func TestPageDeepClaims(t *testing.T) {
	ids := []string{"a", "b"}
	clone := CopyPageIDs(ids)
	clone[0] = "x"
	if ids[0] != "a" {
		t.Fatal("page ids aliased")
	}
	start, end := PageBounds(10, 2, 3)
	if start != 2 || end != 5 {
		t.Fatalf("bounds=%d,%d", start, end)
	}
}
