package cujprocessor

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/processor/processorhelper"
)

const (
	defaultVal = "request.n"
	// this is the name used to refer to the connector in the config.yaml
	typeStr = "cujprocessor"
)

// NewFactory creates a factory for example connector.
func NewFactory() processor.Factory {
	// OpenTelemetry connector factory to make a factory for connectors
	cType, err := component.NewType(typeStr)
	if err != nil {
		return nil
	}

	return processor.NewFactory(
		cType,
		createDefaultConfig,
		processor.WithTraces(createProcessor, component.StabilityLevelStable))
}

func createDefaultConfig() component.Config {
	return &Config{
		AttributeName: defaultVal,
	}
}

// createTracesToMetricsConnector defines the consumer type of the connector
// We want to consume traces and export metrics, therefore, define nextConsumer as metrics, since consumer is the next component in the pipeline
func createProcessor(ctx context.Context, set processor.Settings, cfg component.Config, nextConsumer consumer.Traces) (processor.Traces, error) {
	c, err := newProcessor(set.Logger, cfg)
	if err != nil {
		return nil, err
	}
	return processorhelper.NewTracesProcessor(ctx, set, cfg, nextConsumer, c.processTraces)
}
