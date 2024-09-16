default: build

build:
	~/go/bin/builder --config builder-config.yaml

run:
	./dist/otel-cujprocessor --config run-config/otel-config.yaml