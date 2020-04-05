# Getting Started

Due to golang registry still cache jaeger exporter v1.0.0 that have a code does not supports jaeger exporter new function. To workaround for goprivate it should be point directly, not sure that why it force install jaeger exporter v1.0.0 that the version does not support with init new provider.

**Issue** <https://github.com/open-telemetry/opentelemetry-go/issues/470>

**It should be fixed by running the command as below.**

```sh
go get -v go.opentelemetry.io/otel/exporter/trace/jaeger@v0.2.1
```

# Run

```sh
docker-compose up --build
```

# Example

#### Basic Parent-Children span

example/connect_the_span ![screenshot](https://raw.githubusercontent.com/iamgoangle/opentelemetry-jaeger-exporter/master/example/connect_the_span/screenshot.png)

```sh
go run .
```

open <http://localhost:16686/> to see all trace.