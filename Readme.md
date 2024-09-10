curl --proto '=https' --tlsv1.2 -fL -o ocb \
    https://github.com/open-telemetry/opentelemetry-collector-releases/releases/download/cmd%2Fbuilder%2Fv0.108.1/ocb_0.108.1_darwin_arm64
chmod +x ocb

~/go/bin/builder --config builder-config.yaml