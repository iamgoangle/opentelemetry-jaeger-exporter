package basic_kafka

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/iamgoangle/opentelemetry-jaeger-exporter/internal/otel"
)

type Producer struct {
	Tracer otel.Tracer
}

type KafkaMsg struct {
	Body    string `json:"body"`
	TraceID string `json:"traceId"`
	SpanID  string `json:"spanId"`
}

func NewProducer(t otel.Tracer) *Producer {
	return &Producer{
		Tracer: t,
	}
}

func (p *Producer) Produce(ctx context.Context, body string) ([]byte, error) {
	thisCtx, span := p.Tracer.TracerStart(ctx, "produce")
	defer span.End()

	time.Sleep(5 * time.Second)

	traceId := p.Tracer.TraceID(thisCtx)
	spanId := p.Tracer.SpanID(thisCtx)

	kafkaMsg := &KafkaMsg{
		Body:    body,
		TraceID: traceId,
		SpanID:  spanId,
	}

	b, err := json.Marshal(kafkaMsg)
	if err != nil {
		return nil, err
	}

	log.Printf("produce %+v\n", kafkaMsg)
	log.Println("finish producer")

	return b, nil
}
