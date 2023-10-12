package pcontext

import (
	"context"

	"github.com/goinbox/golog"
)

type Context interface {
	context.Context

	Logger() golog.Logger
}

type simpleContext struct {
	context.Context

	logger golog.Logger
}

func NewSimpleContext(ctx context.Context, logger golog.Logger) Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return &simpleContext{
		Context: ctx,
		logger:  logger,
	}
}

func (s *simpleContext) Logger() golog.Logger {
	return s.logger
}
