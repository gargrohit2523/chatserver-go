package trace

import (
	"fmt"
	"io"
)

// Tracer is the interface that describes an object capable of
// tracing events throughout code.
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

func New(w io.Writer) *tracer {
	return &tracer{out: w}
}

func (t *tracer) Trace(msg ...interface{}) {
	fmt.Fprint(t.out, msg...)
	fmt.Fprint(t.out, "\n")
}
