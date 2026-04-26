package model

import (
	"testing"
)

func TestCheckStatus(t *testing.T) {
	t.Run("successful check for 'DONE' input", func(t *testing.T) {
		err := checkStatus("DONE")

		assertNoError(t, err)
	})

	t.Run("successful check for 'IN_PROGRESS' input", func(t *testing.T) {
		err := checkStatus("IN_PROGRESS")

		assertNoError(t, err)
	})

	t.Run("successful check for 'TODO' input", func(t *testing.T) {
		err := checkStatus("TODO")

		assertNoError(t, err)
	})

	t.Run("successful check for non-valid input", func(t *testing.T) {
		err := checkStatus("invalid_input")

		assertError(t, err)
	})
}

func TestSanitizeString(t *testing.T) {
	type testCase struct {
		name     string
		title    string
		expected string
	}

	testCases := []testCase{
		{
			name:     "sucessful title sanitizing",
			title:    "task title number 1   ",
			expected: "task title number 1",
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got := sanitizeString(test.title)

			if test.expected != got {
				t.Errorf("got %s, expected %s, given %s", got, test.expected, test.title)
			}
		})
	}
}

func TestIsTitleEmpty(t *testing.T) {
	t.Run("title is empty", func(t *testing.T) {
		got := isTitleEmpty("")

		assertError(t, got)
	})

	t.Run("title is not empty", func(t *testing.T) {
		got := isTitleEmpty("not empty")

		assertNoError(t, got)
	})
}

// Helper function to assert no error occured during testing
func assertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("got an error, excpected none: %v", err)
	}
}

// Helper function to assert an error occured during testing
func assertError(t *testing.T, err error, params ...any) {
	t.Helper()

	if err == nil {
		t.Errorf("got no error, expected one given the following parameters: %v", params...)
	}
}
