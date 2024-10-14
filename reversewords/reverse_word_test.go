package reversewords

import "testing"

func TestReverseWord(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Test 1", "Hello World", "olleH dlroW"},
		{"Test 2", "This is an example!", "sihT si na !elpmaxe"},
		{"Test 3", "double  spaces", "elbuod  secaps"},
		{"Test 4", "This is a test", "sihT si a tset"},
	}
	t.Run("Test Reverse Words", func(t *testing.T) {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				actual := ReverseWords(tt.input)
				if actual != tt.expected {
					t.Errorf("Expected %s but got %s", tt.expected, actual)
				}
			})
		}
	})
}
