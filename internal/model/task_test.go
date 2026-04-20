package model

import (
	"testing"
)

func TestCheckStatus(t *testing.T) {
	t.Run("successful check for 'done' input", func(t *testing.T) {
		err := checkStatus("done")

		assertNoError(t, err)
	})

	t.Run("successful check for 'in_progress' input", func(t *testing.T) {
		err := checkStatus("in_progress")

		assertNoError(t, err)
	})

	t.Run("successful check for 'todo' input", func(t *testing.T) {
		err := checkStatus("todo")

		assertNoError(t, err)
	})

	t.Run("successful check for non-valid input", func(t *testing.T) {
		err := checkStatus("invalid_input")

		assertError(t, err)
	})
}

// Helper function to assert no error occured during testing
func assertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("got an error, excpected none: %v", err)
	}
}

func assertError(t *testing.T, err error, params ...any) {
	t.Helper()

	if err == nil {
		t.Errorf("got no error, expected one given the following parameters: %v", params...)
	}
}
