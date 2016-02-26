package tablebased

import (
	"testing"

	"github.com/jaeyeom/gogo/seq"
)

func TestFibNth(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{0, 0},
		{5, 5},
		{6, 8},
	}
	for _, c := range cases {
		got := seq.FibNth(c.in)
		if got != c.want {
			t.Errorf("Fib(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
