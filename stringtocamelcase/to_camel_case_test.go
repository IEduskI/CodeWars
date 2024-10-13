package stringtocamelcase

import "testing"

func TestToCamelCase(t *testing.T) {
	s := "text-to-camel_case"

	want := "textToCamelCase"

	got := ToCamelCase(s)

	if got != want {
		t.Errorf("error got is different to want - got: %s want %s", got, want)
	}
}
