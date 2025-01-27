package twosums

import "testing"

func TestAdd(t *testing.T) {
	got := Sum(8, 2)
	want := 10

	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}
}
