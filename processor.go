package cujprocessor

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

type processorImpl struct {
	config Config
	logger *zap.Logger
}

func newProcessor(logger *zap.Logger, config component.Config) (*processorImpl, error) {
	logger.Info("Building cujprocessor")
	cfg := config.(*Config)

	return &processorImpl{
		config: *cfg,
		logger: logger,
	}, nil
}

func (p *processorImpl) processTraces(_ context.Context, td ptrace.Traces) (ptrace.Traces, error) {
	resourceSpans := td.ResourceSpans()
	for i := 0; i < resourceSpans.Len(); i++ {
		scopeSpans := resourceSpans.At(i).ScopeSpans()
		for j := 0; j < scopeSpans.Len(); j++ {
			spans := scopeSpans.At(j).Spans()
			for k := 0; k < spans.Len(); k++ {
				span := spans.At(k)
				// Ajouter une clé statique à chaque span
				attributes := span.Attributes()
				attributes.PutStr("custom_key", "static_value")
			}
		}
	}
	return td, nil
}
