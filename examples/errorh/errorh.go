// Example for the new error handling.
package errorh

import "golang.org/x/xerrors"

var (
	ErrMyError     = xerrors.New("my error")
	ErrCannotClose = xerrors.New("error while closing")
)

func myFunc() error {
	return ErrMyError
}

func f() error {
	err := myFunc()
	if err != nil {
		return xerrors.Errorf("error while calling myFunc: %w", err)
	}
	return nil
}

func g() error {
	err := f()
	if err != nil {
		return xerrors.Errorf("error while calling f: %w", err)
	}
	return nil
}

func c() error {
	return ErrCannotClose
}

func h() (err error) {
	defer func() {
		e := c()
		if e != nil {
			// Choose to lose wrapping information of ErrCannotClose.
			err = xerrors.Errorf("%v: %w", e, err)
		}
	}()
	err = g()
	if err != nil {
		return xerrors.Errorf("error while calling g: %w", err)
	}
	return nil
}
