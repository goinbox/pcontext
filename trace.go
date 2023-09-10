package pcontext

import (
	"go.opentelemetry.io/otel/trace"
)

type StartTraceFunc[T Context] func(ctx T, spanName string, opts ...trace.SpanStartOption) (T, trace.Span)
