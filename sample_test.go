package main

import "testing"

func TestHelloname(t *testing.T) {
	result := HelloName("john doe")

	if result != "Hello john doe" {
		t.Fatal("return error")
	}
}
