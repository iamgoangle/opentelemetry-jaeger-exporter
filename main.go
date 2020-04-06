package main

import (
	"context"
	"os"
	"time"

	"github.com/iamgoangle/opentelemetry-jaeger-exporter/kafka"

	"go.opentelemetry.io/otel/api/global"

	"github.com/iamgoangle/opentelemetry-jaeger-exporter/internal/otel"
)

func main() {
	initTrace := otel.InitTracer(&otel.Config{
		Service:        "distribute-trace-app",
		ThriftEndpoint: os.Getenv("THRIFT_ENDPOINT"),
	})
	defer initTrace()

	tracer := otel.NewTracer()
	ctx := context.Background()

	tr := global.TraceProvider().Tracer("trace-main")
	ctx, span := tr.Start(ctx, "main")

	// produce app
	// actually, this should be a micro service run in isolate container
	broker := []string{os.Getenv("KAFKA_BROKER")}
	producer := kafka.NewProducer(os.Getenv("KAFKA_TOPIC"), broker, tracer)
	producer.Produce(ctx, "test tracing ja")

	span.End()

	time.Sleep(5 * time.Second)

	// consumer app
	// actually, this should be a micro service run in isolate container
	consumer := kafka.NewConsumer(os.Getenv("KAFKA_TOPIC"), broker, tracer)
	consumer.Consume()
}

// func main() {
// 	initTrace := otel.InitTracer(&otel.Config{
// 		Service:        "producer-app",
// 		ThriftEndpoint: os.Getenv("THRIFT_ENDPOINT"),
// 	})
// 	defer initTrace()
//
// 	tracer := otel.NewTracer()
// 	ctx := context.Background()
//
// 	tr := global.TraceProvider().Tracer("trace-main")
// 	ctx, span := tr.Start(ctx, "main")
//
// 	// producer
// 	producer := basic_kafka.NewProducer(tracer)
// 	body, err := producer.Produce(ctx, "test test test")
// 	if err != nil {
// 		log.Panic(err)
// 	}
//
// 	span.End()
//
// 	// consumer
// 	tracer = otel.NewTracer()
// 	consumer := basic_kafka.NewConsumer(tracer)
// 	err = consumer.Consume(body)
// 	if err != nil {
// 		log.Panic(err)
// 	}
//
// 	// apiTest := &api.TestHandler{
// 	// 	Tracer: tracer,
// 	// }
// 	// apiTest.Handler(ctx)
// }
