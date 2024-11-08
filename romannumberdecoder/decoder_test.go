package romannumberdecoder

import "testing"

func TestDecode(t *testing.T) {
	tests := []struct {
		roman string
		want  int
	}{
		{"XXI", 21},
		{"I", 1},
		{"IV", 4},
		{"MMVIII", 2008},
		{"MDCLXVI", 1666},
	}

	for _, tt := range tests {
		t.Run(tt.roman, func(t *testing.T) {
			got := Decode(tt.roman)
			if got != tt.want {
				t.Errorf("Decode(%s) = %d; want %d", tt.roman, got, tt.want)
			}
		})
	}
}
