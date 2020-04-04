package main

import (
	"context"
	"net/http"
	"time"

	"google.golang.org/grpc/codes"

	"go.opentelemetry.io/otel/api/global"

	"github.com/iamgoangle/opentelemetry-jaeger-exporter/internal/otel"
)

type TestHandler struct {
	Tracer otel.Tracer
}

func main() {
	initTrace := otel.InitTracer(&otel.Config{
		Service:        "my-app",
		ThriftEndpoint: "http://localhost:14268/api/traces",
	})
	defer initTrace()

	ctx := context.Background()

	tr := global.TraceProvider().Tracer("main")
	ctx, span := tr.Start(ctx, "main-span")
	defer span.End()

	tracer := otel.NewTracer()

	test := &TestHandler{
		Tracer: tracer,
	}
	test.handler(ctx)
}

func (t *TestHandler) handler(ctx context.Context) {
	spanName := "handler"
	thisCtx, span := t.Tracer.TracerStart(ctx, spanName)
	defer span.End()

	time.Sleep(10 * time.Second)

	t.Tracer.SetIntAttribute(thisCtx, "http.code", http.StatusOK)
	span.SetStatus(codes.OK)

	t.service(thisCtx)
}

func (t *TestHandler) service(ctx context.Context) {
	spanName := "service"
	thisCtx, span := t.Tracer.TracerStart(ctx, spanName)
	defer span.End()

	time.Sleep(10 * time.Second)

	t.repository(thisCtx)
}

func (t *TestHandler) repository(ctx context.Context) {
	spanName := "repository"
	_, span := t.Tracer.TracerStart(ctx, spanName)
	defer span.End()

	time.Sleep(10 * time.Second)
}

// func sayHello(ctx context.Context, name string) {
// 	trace := global.TraceProvider().Tracer("sayHello")
// 	spanName := fmt.Sprintf("Hello, %s", name)
//
// 	ctx, span := trace.Start(ctx, spanName)
// 	defer span.End()
//
// 	time.Sleep(10 * time.Second)
//
// 	sayProfileByName(ctx, name)
// }
//
// func sayProfileByName(ctx context.Context, name string) {
// 	trace := global.TraceProvider().Tracer("sayProfileByName")
// 	spanName := fmt.Sprintf("Profile, %s", name)
//
// 	_, span := trace.Start(ctx, spanName)
// 	defer span.End()
//
// 	time.Sleep(10 * time.Second)
// }
//
// func asyncJob(ctx context.Context) {
// 	trace := global.TraceProvider().Tracer("asyncJob")
// 	spanName := fmt.Sprintf("Async Job")
//
// 	_, span := trace.Start(ctx, spanName)
// 	defer span.End()
//
// 	time.Sleep(2 * time.Second)
// }
//
// func sayGoodBye(ctx context.Context, name string) {
// 	trace := global.TraceProvider().Tracer("sayGoodBye")
// 	spanName := fmt.Sprintf("Bye, %s", name)
//
// 	_, span := trace.Start(ctx, spanName)
// 	defer span.End()
//
// 	time.Sleep(10 * time.Second)
// }
