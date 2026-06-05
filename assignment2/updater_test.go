package main

import "testing"

func TestUpdateValue(t *testing.T) {
	if got := UpdateValue(10, 5); got != 15 {
		t.Fatalf("UpdateValue(10, 5) = %d, want 15", got)
	}
}
