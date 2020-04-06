package kafka

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/Shopify/sarama"

	"github.com/iamgoangle/opentelemetry-jaeger-exporter/internal/otel"
)

type Producer struct {
	Topic string
	Broker []string
	Trace otel.Tracer
}

type KafkaTracing struct {
	Body string `json:"body"`
	TraceID string `json:"traceId"`
	SpanID string `json:"spanId"`
}

func NewProducer(topic string, bk []string, t otel.Tracer) *Producer {
	return &Producer{
		Topic: topic,
		Broker: bk,
		Trace: t,
	}
}

func (p *Producer) Produce(ctx context.Context, body string) {
	thisCtx, span := p.Trace.TracerStart(ctx, "Kafka/Produce")
	defer span.End()

	producer, err := sarama.NewSyncProducer(p.Broker, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// append trace to kafka body
	traceId := p.Trace.TraceID(thisCtx)
	spanId := p.Trace.SpanID(thisCtx)
	log.Println("Produce TraceID: ", traceId)
	log.Println("Produce SpanID: ", spanId)

	bodyWithTrace := &KafkaTracing{
		Body: body,
		TraceID: traceId,
		SpanID: spanId,
	}

	b, _ := json.Marshal(bodyWithTrace)

	msg := &sarama.ProducerMessage{
		Topic:     p.Topic,
		Value:     sarama.ByteEncoder(b),
		Timestamp: time.Now(),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
	} else {
		log.Printf("Produce message sent to partition %d at offset %d\n", partition, offset)
		log.Printf("Produce Body => %+v\n\n", string(body))
	}
}
