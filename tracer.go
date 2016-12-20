package trace

import (
	"io"
	"fmt"
)

type tracer struct {
	out io.Writer
}
func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

type Tracer interface {
	Trace(...interface{})
}
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type nilTracer struct{}
func (t *nilTracer) Trace(a ...interface{}) {}

func Off() Tracer {
	return &nilTracer{}
}
