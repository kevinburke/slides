package main

import "testing"

func TestAdd(t *testing.T) {
	if got := add('M', 'E'); got != 'Q' {
		t.Errorf("expected add('M', 'E') to equal 'Q', got %q", got)
	}
}
