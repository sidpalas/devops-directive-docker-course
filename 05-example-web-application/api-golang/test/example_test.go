package test

import "testing"

func TestOneEqualsOne(t *testing.T) {
	val := 1
	if val != 1 {
		t.Errorf("1 != %d; want 1", val)
	}
}
