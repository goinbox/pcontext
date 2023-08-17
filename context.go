package pcontext

import (
	"context"

	"github.com/goinbox/golog"

	"go.opentelemetry.io/otel/trace"
)

type Context interface {
	context.Context

	Logger() golog.Logger
	BeginTrace(spanName string, opts ...trace.SpanStartOption) (Context, trace.Span)
}
