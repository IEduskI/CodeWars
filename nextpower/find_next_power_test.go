package nextpower

import "testing"

func TestFindNextPower(t *testing.T) {
	tests := []struct {
		val, pow, want int
	}{
		{12385, 3, 13824},
		{1245678, 5, 1419857},
	}
	for _, test := range tests {
		if got := FindNextPower(test.val, test.pow); got != test.want {
			t.Errorf("FindNextPower(%d, %d) = %d; want %d", test.val, test.pow, got, test.want)
		}
	}
}
