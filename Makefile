default: build

tidy:
	cd src && go mod tidy

build:
	~/go/bin/builder --config builder-config.yaml

run:
	./dist/otel-cujprocessor --config run-config/otel-config.yaml