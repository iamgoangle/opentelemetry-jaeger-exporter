package main

import (
	"context"
	"log"

	"go.opentelemetry.io/otel/api/global"

	"github.com/iamgoangle/opentelemetry-jaeger-exporter/internal/otel"

	// "github.com/iamgoangle/opentelemetry-jaeger-exporter/api"
	"github.com/iamgoangle/opentelemetry-jaeger-exporter/kafka"
)

func main() {
	initTrace := otel.InitTracer(&otel.Config{
		Service:        "producer-app",
		ThriftEndpoint: "http://localhost:14268/api/traces",
	})
	defer initTrace()

	tracer := otel.NewTracer()
	ctx := context.Background()

	tr := global.TraceProvider().Tracer("trace-main")
	ctx, span := tr.Start(ctx, "main")

	// producer
	producer := kafka.NewProducer(tracer)
	body, err := producer.Produce(ctx, "test test test")
	if err != nil {
		log.Panic(err)
	}

	span.End()

	// consumer
	tracer = otel.NewTracer()
	consumer := kafka.NewConsumer(tracer)
	err = consumer.Consume(body)
	if err != nil {
		log.Panic(err)
	}

	// apiTest := &api.TestHandler{
	// 	Tracer: tracer,
	// }
	// apiTest.Handler(ctx)
}
