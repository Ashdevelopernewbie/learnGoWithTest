package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	expectedLang := "english"
	t.Run("empty name with empty language", func(t *testing.T) {
		got, language, err := Hello("", "")
		expected := "Hello, world! noob"

		checkErrorMessage(t, err)
		if language != expectedLang {
			assertCorrectMessage(t, language, expectedLang)
		}
		assertCorrectMessage(t, got, expected)
	})
	t.Run("name: world", func(t *testing.T) {
		got, language, err := Hello("world", "")
		expected := "Hello, world! noob"

		checkErrorMessage(t, err)
		if language != expectedLang {
			assertCorrectMessage(t, language, expectedLang)
		}
		assertCorrectMessage(t, got, expected)
	})
	t.Run("name: anmol", func(t *testing.T) {
		got, language, err := Hello("anmol", "")
		expected := "Hello, anmol! notnoob"

		checkErrorMessage(t, err)
		if language != expectedLang {
			assertCorrectMessage(t, language, expectedLang)
		}
		assertCorrectMessage(t, got, expected)
	})

}

func checkErrorMessage(t testing.TB, err error) {
	if err != nil {
		t.Fatalf("no error expected but found: %v", err)
	}
}

func assertCorrectMessage(t testing.TB, got, expected string) {
	t.Helper()
	if got != expected {
		t.Errorf("\ngot:      %v\nexpected: %v", got, expected)
	} else {
		t.Logf("\ngot:      %v\nexpected: %v", got, expected)
	}
}
