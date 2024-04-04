package main

import "testing"

func TestReverseString(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"Ерасыл", "лысарЕ"},
		{"", ""},
		{"12345", "54321"},
		{"Привет, друг!", "!гурд ,тевирП"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := ReverseString(tc.input)
			if result != tc.expected {
				t.Errorf("Ожидалось :  %q, фактический результат %q", tc.expected, result)
			}
		})
	}
}
