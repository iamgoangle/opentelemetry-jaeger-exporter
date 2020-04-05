package kafka

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel/api/trace"

	"encoding/json"

	"github.com/iamgoangle/opentelemetry-jaeger-exporter/internal/otel"
	"go.opentelemetry.io/otel/api/core"
)

type Consumer struct {
	Tracer otel.Tracer
}

func NewConsumer(t otel.Tracer) *Consumer {
	return &Consumer{
		Tracer: t,
	}
}

func (c *Consumer) Consume(msg []byte) error {
	initTrace := otel.InitTracer(&otel.Config{
		Service:        "consumer-app",
		ThriftEndpoint: "http://localhost:14268/api/traces",
	})
	defer initTrace()

	var kafkaBody KafkaMsg
	err := json.Unmarshal(msg, &kafkaBody)
	if err != nil {
		return err
	}

	log.Printf("consumer body %+v \n", kafkaBody)

	traceId, _ := core.TraceIDFromHex(kafkaBody.TraceID)
	spanId, _ := core.SpanIDFromHex(kafkaBody.SpanID)
	sc := core.SpanContext{
		TraceID:    traceId,
		SpanID:     spanId,
		TraceFlags: 0x0,
	}

	_, span := c.Tracer.TracerStart(trace.ContextWithRemoteSpanContext(context.Background(), sc), "consume")
	defer span.End()

	time.Sleep(5 * time.Second)

	log.Println("finish consumer")

	return nil
}
