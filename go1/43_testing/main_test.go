package main_test

import (
	"testing"
)

// running test: go test main_test.go
// more details: see guide -> TOOLS

func TestSuccess(t *testing.T) {
	sum := 2 + 2
	if sum != 4 {
		t.Errorf("sum is not correct")
	}
}

func TestError(t *testing.T) {
	t.Error("fatal error")
}
