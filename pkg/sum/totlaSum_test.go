package sum

import (
	"testing"
)

func TestInefficientSum(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{input: 10, expected: 55},
		{input: 100, expected: 5050},
		{input: 9876, expected: 48772626},
		{input: 123456789, expected: 7620789436823655},
	}

	for _, test := range tests {
		result := InefficientSum(test.input)
		if result != test.expected {
			t.Errorf("InefficientSum(%d) = %d; want %d", test.input, result, test.expected)
		}
	}
}

func TestEfficientSum(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{input: 10, expected: 55},
		{input: 100, expected: 5050},
		{input: 9876, expected: 48772626},
		{input: 123456789, expected: 7620789436823655},
	}

	for _, test := range tests {
		result := EfficientSum(test.input)
		if result != test.expected {
			t.Errorf("InefficientSum(%d) = %d; want %d", test.input, result, test.expected)
		}
	}
}

func BenchmarkInefficientSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InefficientSum(i * 1000)
	}
}

func BenchmarkEfficientSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EfficientSum(i * 1000)
	}
}
