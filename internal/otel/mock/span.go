package mock

import (
	"context"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/trace"
)

// MockSpan represent mock opentelemetry
/**
	Example:

	It("should call handler", func() {
		tr := global.TraceProvider().Tracer("mockTracer")
		_, span := tr.Start(mockCtx, "mockSpan")
		defer span.End()

		mockTraceID := "58e32b2438efda4d90db7b0ce0db9371"
		mockSpanID := "8fc213b9e9b84a98"
		mockTracer.EXPECT().TracerStart(mockCtx, "handler").Return(mockCtx, span).Times(1)
		mockTracer.EXPECT().SetIntAttribute(mockCtx, gomock.Any(), gomock.Any()).Times(1)
		mockTracer.EXPECT().TraceID(mockCtx).Return(mockTraceID).Times(1)
		mockTracer.EXPECT().SpanID(mockCtx).Return(mockSpanID).Times(1)

		mockAPI := &TestHandler{
			Tracer: mockTracer,
		}

		mockAPI.Handler(mockCtx)
	})
 */
func MockSpan(ctx context.Context) (context.Context, trace.Span) {
	tr := global.TraceProvider().Tracer("mockTracer")
	thisCtx, span := tr.Start(ctx, "mockSpan")
	defer span.End()

	return thisCtx, span
}
