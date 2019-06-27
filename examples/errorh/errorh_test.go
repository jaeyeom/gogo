package errorh

import (
	"fmt"

	"golang.org/x/xerrors"
)

func Example_wrapping() {
	err := h()
	// fmt.Printf("%+v\n", err) // will produce the trace
	fmt.Println(xerrors.Is(err, ErrMyError))
	fmt.Println(xerrors.Is(err, ErrCannotClose))
	for ; err != nil; err = xerrors.Unwrap(err) {
		fmt.Println(err)
	}
	// Output Example (for %+v trace):
	// error while closing:
	// github.com/jaeyeom/gogo/examples/errorh.h.func1
	// /home/jaeyeom/go/src/github.com/jaeyeom/gogo/examples/errorh/errorh.go:40
	// - error while calling g:
	// github.com/jaeyeom/gogo/examples/errorh.h
	// /home/jaeyeom/go/src/github.com/jaeyeom/gogo/examples/errorh/errorh.go:45
	// - error while calling f:
	// github.com/jaeyeom/gogo/examples/errorh.g
	// /home/jaeyeom/go/src/github.com/jaeyeom/gogo/examples/errorh/errorh.go:26
	// - error while calling myFunc:
	// github.com/jaeyeom/gogo/examples/errorh.f
	// /home/jaeyeom/go/src/github.com/jaeyeom/gogo/examples/errorh/errorh.go:18
	// - my error:
	// github.com/jaeyeom/gogo/examples/errorh.init.ializers
	// /home/jaeyeom/go/src/github.com/jaeyeom/gogo/examples/errorh/errorh.go:7

	// Output:
	// true
	// false
	// error while closing: error while calling g: error while calling f: error while calling myFunc: my error
	// error while calling g: error while calling f: error while calling myFunc: my error
	// error while calling f: error while calling myFunc: my error
	// error while calling myFunc: my error
	// my error
}
