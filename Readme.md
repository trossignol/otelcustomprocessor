# Custom OpenTelemetry Processor

This project is an example of a custom OpenTelemetry processor. It demonstrates how to create a custom OpenTelemetry processor in Go and how to integrate it with the OpenTelemetry ecosystem. This processor is intended solely for demonstration purposes.

In `src` directory, you can find 3 `go` files:
* `config.go` used to specify the processor's configuration parameters
* `factory.go` containing the instantiation of the processor
* `processor.go` which is the "real" processor code (especially the `processTraces` method)

## Prerequisites

Before you begin, ensure you have the following installed on your machine:

- [Make](https://www.gnu.org/software/make/) (for build and run tasks)
- [Go](https://golang.org/doc/install)
- [OTEL Builder](https://opentelemetry.io/docs/collector/custom-collector/)

## Installation

Clone this repository:

```bash
git clone https://github.com/your-username/custom-otel-processor.git
cd custom-otel-processor
```

## Build

To build the collector, use the `build` task in the `Makefile`:

```bash
make build
```

## Run

To run the collector, use the `run` task in the `Makefile`:

```bash
make run
```

The result is an `otel-cujprocessor` binary available in the `dist` directory.