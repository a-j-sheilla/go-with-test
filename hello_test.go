package main

import (
	//"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, World"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	} 
}
