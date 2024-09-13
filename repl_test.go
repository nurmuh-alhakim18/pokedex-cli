package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:  "lower case input",
			input: "hello world",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			name:  "upper case input",
			input: "HELLO WORLD",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := cleanInput(tc.input)
			if len(actual) != len(tc.expected) {
				t.Errorf("The lengths are not equal: %v vs %v", len(actual), len(tc.expected))
				return
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("%v does not equal to %v", actual, tc.expected)
				return
			}
		})
	}
}
