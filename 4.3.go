package main

import "testing"

func testHello(t *testing.T) {
	got := Hello, World()
	want := "Hello, world"

	if got != want {
		t.Error("got %q want %q", got, want)
	}
}