package testhelpers

import (
	"strings"
	"testing"
)

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v but wanted %v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("got %v but wanted %v", got, want)
	}
}

func AssertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("An error has occurred: %v", err)
	}
}

func AssertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Errorf("We were expecting an error: %v", err)
	}
}

func AssertContains(t *testing.T, returned, included string) {
	t.Helper()
	if !strings.Contains(returned, included) {
		t.Errorf("The string %s does not contain %s", returned, included)
	}
}

func AssertTrue(t *testing.T, expected any) {
	t.Helper()
	if expected == false {
		t.Errorf("Was expecting %v to be true but it is %v", expected, false)
	}
}
