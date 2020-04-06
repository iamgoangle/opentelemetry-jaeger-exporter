package kafka

import (
	"context"
	"encoding/json"
	"go.opentelemetry.io/otel/api/core"
	"go.opentelemetry.io/otel/api/trace"
	"log"
	"os"

	"github.com/iamgoangle/opentelemetry-jaeger-exporter/internal/otel"

	"github.com/Shopify/sarama"
)

type exampleConsumerGroupHandler struct {
	Trace otel.Tracer
}

func (h *exampleConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (h *exampleConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h *exampleConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	initTrace := otel.InitTracer(&otel.Config{
		Service:        "consumer-app",
		ThriftEndpoint: os.Getenv("THRIFT_ENDPOINT"),
	})

	defer initTrace()

	for msg := range claim.Messages() {
		log.Printf("Consumer Message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
		log.Printf("Consumer Message => %+v\n", string(msg.Value))

		// consumer get message from kafka topic
		// and extract the traceId, spanId
		// for continue the trace
		var kafkaBody KafkaTracing
		err := json.Unmarshal(msg.Value, &kafkaBody)
		if err != nil {
			panic(err)
		}

		traceId, _ := core.TraceIDFromHex(kafkaBody.TraceID)
		spanId, _ := core.SpanIDFromHex(kafkaBody.SpanID)
		sc := core.SpanContext{
			TraceID:    traceId,
			SpanID:     spanId,
			TraceFlags: 0x0,
		}

		_, span := h.Trace.TracerStart(trace.ContextWithRemoteSpanContext(context.Background(), sc), "Kafka/ConsumeClaim")
		log.Println("Consumer TraceID: ", span.SpanContext().TraceIDString())
		log.Println("Consumer SpanID: ", span.SpanContext().SpanIDString())

		// mark offset done
		sess.MarkMessage(msg, "")

		span.End()
	}
	return nil
}

type Consumer struct {
	Topic  string
	Broker []string
	Tracer otel.Tracer
}

func NewConsumer(topic string, bk []string, t otel.Tracer) *Consumer {
	return &Consumer{
		Topic:  topic,
		Broker: bk,
		Tracer: t,
	}
}

func (c *Consumer) Consume() {
	// Init config, specify appropriate version
	config := sarama.NewConfig()
	config.Version = sarama.V1_0_0_0
	config.Consumer.Return.Errors = true

	// Start with a client
	client, err := sarama.NewClient(c.Broker, config)
	if err != nil {
		panic(err)
	}
	defer func() { _ = client.Close() }()

	// Start a new consumer group
	group, err := sarama.NewConsumerGroupFromClient(os.Getenv("CONSUMER_GROUP"), client)
	if err != nil {
		panic(err)
	}
	defer func() { _ = group.Close() }()

	// Track errors
	go func() {
		for err := range group.Errors() {
			log.Println("ERROR", err)
		}
	}()

	// Iterate over consumer sessions.
	ctx := context.Background()
	for {
		topics := []string{c.Topic}
		handler := &exampleConsumerGroupHandler{
			Trace: c.Tracer,
		}

		err := group.Consume(ctx, topics, handler)
		if err != nil {
			panic(err)
		}
	}
}
