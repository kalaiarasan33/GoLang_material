package main

import "testing"

func TestAdd(t *testing.T) {
	t.Run("Test1", func(t *testing.T) {
		got := Add(10, 20)
		want := 30
		if got != want {
			t.Errorf("Got: %v - want: %v", got, want)
		}
	})
	t.Run("Test2", func(t *testing.T) {
		got := Add(20, 20)
		want := 40
		if got != want {
			t.Errorf("Got: %v - want: %v", got, want)
		}

	})
}
