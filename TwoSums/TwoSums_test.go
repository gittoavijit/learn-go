package twosums

import "testing"

func TestAdd(t *testing.T) {
	sum := 4 + 6
	got := sum
	want := 10

	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}
}
