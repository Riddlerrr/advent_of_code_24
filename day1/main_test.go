package main

import (
	"slices"
	"testing"
)

func TestReadFile(t *testing.T) {
	t.Run("read from file", func(t *testing.T) {
		got := ReadFile("test_example.txt")
		want := []string{
			"3   4",
			"4   3",
			"2   5",
			"1   3",
			"3   9",
			"3   3",
		}

		assertCorrectFileContent(t, got, want)
	})
}

func TestParseLine(t *testing.T) {
	tests := []struct {
		input         string
		expectedLeft  int
		expectedRight int
	}{
		{"1   2", 1, 2},
		{"10   20", 10, 20},
		{"100   200", 100, 200},
		{"", 0, 0},
		{"1 2", 0, 0},
		{"1   ", 0, 0},
		{"   2", 0, 0},
		{"abc   def", 0, 0},
	}

	for _, test := range tests {
		left, right := parseLine(test.input)
		if left != test.expectedLeft || right != test.expectedRight {
			t.Errorf("parseLine(%q) = (%d, %d); want (%d, %d)", test.input, left, right, test.expectedLeft, test.expectedRight)
		}
	}
}

func assertCorrectFileContent(t testing.TB, got, want []string) {
	t.Helper()
	if slices.Compare(got, want) != 0 {
		t.Errorf("got %q want %q", got, want)
	}
}
