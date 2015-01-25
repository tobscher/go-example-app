package main

import "testing"

func TestSanity(t *testing.T) {
	if true == false {
		t.Error("True should not be false.")
	}
}
