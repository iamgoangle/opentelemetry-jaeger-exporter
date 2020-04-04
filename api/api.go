package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/iamgoangle/opentelemetry-jaeger-exporter/internal/otel"
	"google.golang.org/grpc/codes"
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

	// t.Service(thisCtx)
}

func (t *TestHandler) Service(ctx context.Context) {
	spanName := "service"
	_, span := t.Tracer.TracerStart(ctx, spanName)
	defer span.End()

	time.Sleep(10 * time.Second)

	// t.Repository(thisCtx)
}

func (t *TestHandler) Repository(ctx context.Context) {
	spanName := "repository"
	_, span := t.Tracer.TracerStart(ctx, spanName)
	defer span.End()

	time.Sleep(10 * time.Second)
}
