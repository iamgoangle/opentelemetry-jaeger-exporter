package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc/codes"

	"github.com/iamgoangle/opentelemetry-jaeger-exporter/internal/otel"
)

type TestHandler struct {
	Tracer otel.Tracer
}

func (t *TestHandler) Handler(ctx context.Context) {
	spanName := "handler"
	thisCtx, span := t.Tracer.TracerStart(ctx, spanName)
	defer span.End()

	time.Sleep(10 * time.Second)

	t.Tracer.SetIntAttribute(thisCtx, "http.code", http.StatusOK)
	span.SetStatus(codes.OK)

	fmt.Println(t.Tracer.TraceID(thisCtx))
	fmt.Println(t.Tracer.SpanID(thisCtx))

	log.Println("finish handler")

	t.Service(thisCtx)
}

func (t *TestHandler) Service(ctx context.Context) {
	spanName := "service"
	thisCtx, span := t.Tracer.TracerStart(ctx, spanName)
	defer span.End()

	time.Sleep(10 * time.Second)

	log.Println("finish service")

	t.Repository(thisCtx)
}

func (t *TestHandler) Repository(ctx context.Context) {
	spanName := "repository"
	_, span := t.Tracer.TracerStart(ctx, spanName)
	defer span.End()

	log.Println("finish repository")

	time.Sleep(10 * time.Second)
}
