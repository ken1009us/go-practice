package main

import (
	"testing"
)

func TestSayHello(t *testing.T) {
	want := "Hello, test!"
	got := SayHello("test")

	if want != got {
		t.Errorf("Wanted %s, got %s", want, got)
	}
}