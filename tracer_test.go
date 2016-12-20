package trace

import (
	"bytes"
	"testing"
)
func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("The return value from new is nil.")
	} else {
		tracer.Trace("Hello, trace package")
		if buf.String() != "Hello, trace package\n" {
			t.Error("The output string is wrong, %s", buf.String())
		}
	}
}

func TestOff(t *testintg.T) {
	var silentTracer Tracer = Off()
	silentTracer.Trace("Data")
}

