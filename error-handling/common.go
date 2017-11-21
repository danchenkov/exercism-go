package erratum

import (
	"io"
)

type TransientError struct {
	err error
}

func (e TransientError) Error() string {
	return e.err.Error()
}

type FrobError struct {
	defrobTag string
	inner     error
}

func (e FrobError) Error() string {
	return e.inner.Error()
}

type Resource interface {
	io.Closer

	Frob(string)
	Defrob(string)
}

type ResourceOpener func() (Resource, error)
