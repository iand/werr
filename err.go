package werr

import (
	"fmt"

	"golang.org/x/xerrors"
)

// Wrap annotates an error with the current stack frame.
func Wrap(err error) error {
	if err == nil {
		return nil
	}
	return &werr{err: err, frame: xerrors.Caller(1)}
}

type werr struct {
	err   error
	frame xerrors.Frame
}

type (
	Printer   = xerrors.Printer
	Formatter = xerrors.Formatter
)

func (e *werr) Unwrap() error              { return e.err }
func (e *werr) Error() string              { return e.err.Error() }
func (e *werr) Format(s fmt.State, v rune) { xerrors.FormatError(e, s, v) }
func (e *werr) FormatError(p Printer) error {
	if ferr, isFormatter := e.err.(Formatter); isFormatter {
		ferr.FormatError(p)
	} else {
		p.Print(e.err.Error())
	}

	if p.Detail() {
		e.frame.Format(p)
	}
	return nil
}
