package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("empty name", func(t *testing.T) {
		_, err := Hello("")
		expectedErr := "name cannot be empty"
		if err == nil || err.Error() != expectedErr {
			t.Fatalf("expected error, got %v", expectedErr)
		} else {
			t.Logf("\ngot:      %v\nexpected: %v", err, expectedErr)
		}
	})
	t.Run("name: world", func(t *testing.T) {
		got, err := Hello("world")
		expected := "Hello, world! noob"

		if err != nil {
			t.Fatalf("no error expected but found: %v", err)
		}
		if got != expected {
			t.Errorf("\ngot:      %v\nexpected: %v", got, expected)
		} else {
			t.Logf("\ngot:      %v\nexpected: %v", got, expected)
		}

	})
	t.Run("name: anmol", func(t *testing.T) {
		got, err := Hello("anmol")
		expected := "Hello, anmol! notnoob"

		if err != nil {
			t.Fatalf("no error expected but found: %v", err)
		}
		if got != expected {
			t.Errorf("\ngot:      %v\nexpected: %v", got, expected)
		} else {
			t.Logf("\ngot:      %v\nexpected: %v", got, expected)
		}
	})
}
