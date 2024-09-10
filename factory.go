package cujprocessor

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/consumer"
)

const (
	defaultVal = "request.n"
	// this is the name used to refer to the connector in the config.yaml
	typeStr = "cujprocessor"
)

// NewFactory creates a factory for example connector.
func NewFactory() connector.Factory {
	// OpenTelemetry connector factory to make a factory for connectors
	cType, err := component.NewType(typeStr)
	if err != nil {
		return nil
	}

	return connector.NewFactory(
		cType,
		createDefaultConfig,
		connector.WithTracesToTraces(createTracesToTracesConnector, component.StabilityLevelAlpha))
	//connector.WithTracesToMetrics(createTracesToMetricsConnector, component.StabilityLevelAlpha))
}

func createDefaultConfig() component.Config {
	return &Config{
		AttributeName: defaultVal,
	}
}

// createTracesToMetricsConnector defines the consumer type of the connector
// We want to consume traces and export metrics, therefore, define nextConsumer as metrics, since consumer is the next component in the pipeline
func createTracesToTracesConnector(ctx context.Context, params connector.Settings, cfg component.Config, nextConsumer consumer.Traces) (connector.Traces, error) {
	c, err := newConnector(params.Logger, cfg)
	if err != nil {
		return nil, err
	}
	c.nextConsumer = nextConsumer
	return c, nil
}
