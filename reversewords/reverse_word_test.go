package reversewords

import "testing"

func TestReverseWord(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Test 1", "Hello World", "olleH dlroW"},
		{"Test 2", "Hello  World", "olleH  dlroW"},
		{"Test 3", "Hello World ", "olleH dlroW "},
		{"Test 4", " Hello World", " olleH dlroW"},
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
