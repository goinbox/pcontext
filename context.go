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

func NewSimpleContext(logger golog.Logger) Context {
	return &simpleContext{
		Context: context.Background(),
		logger:  logger,
	}
}

func (s *simpleContext) Logger() golog.Logger {
	return s.logger
}
