package trace

import (
	"fmt"
	"io"
)

// Tracer is the interface that describes an object capable of
// tracing events throughout code.
type Tracer interface {
	Trace(...any)
}

type tracker struct {
	out io.Writer
}

func (t *tracker) Trace(a ...any) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

func New(w io.Writer) Tracer {
	return &tracker{out: w}
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...any) {}

// Off returns a Tracer that will ingore calls to Trace.
func Off() Tracer {
	return &nilTracer{}
}
