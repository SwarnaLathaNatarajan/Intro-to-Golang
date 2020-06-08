package main

import "testing"

func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Error")
	}
}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-2, 0},
		{999, 1001},
	}
	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("Test Failed : {} inputted, {} expected, {} received", test.input, test.expected, output)
		}
	}
}
