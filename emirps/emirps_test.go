package emirps

import "testing"

func TestEmirps(t *testing.T) {
	tests := []struct {
		n        int
		emirps   int
		largest  int
		sumEmirp int
	}{
		{50, 4, 37, 98},
		{100, 8, 97, 418},
	}
	for _, test := range tests {
		emirps := FindEmirp(test.n)
		if emirps[0] != test.emirps || emirps[1] != test.largest || emirps[2] != test.sumEmirp {
			t.Errorf("FindEmirp(%d) = %d, %d, %d; want %d, %d, %d", test.n, emirps[0], emirps[1], emirps[2], test.emirps, test.largest, test.sumEmirp)
		}
	}
}
