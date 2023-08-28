package pcontext

import (
	"context"

	"github.com/goinbox/golog"
)

type Context interface {
	context.Context

	TraceID() string
	Logger() golog.Logger
}

type simpleContext struct {
	context.Context

	tid    string
	logger golog.Logger
}

func NewSimpleContext(tid string, logger golog.Logger) Context {
	return &simpleContext{
		Context: context.Background(),
		tid:     tid,
		logger:  logger,
	}
}

func (s *simpleContext) TraceID() string {
	return s.tid
}

func (s *simpleContext) Logger() golog.Logger {
	return s.logger
}
