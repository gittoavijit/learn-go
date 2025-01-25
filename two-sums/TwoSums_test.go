package TwoSums

import "testing"

func TestAdd(t *testing.T) {
	got := sum(7, 3)
	want := 10

	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}
}
