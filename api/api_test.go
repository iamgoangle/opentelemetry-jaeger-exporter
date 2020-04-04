package api_test

import (
	"context"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"

	mockOtel "github.com/iamgoangle/opentelemetry-jaeger-exporter/internal/otel/mock"

	. "github.com/iamgoangle/opentelemetry-jaeger-exporter/api"
)

var _ = Describe("ConnectTheSpan", func() {
	var (
		ctrl        *gomock.Controller
		mockTracer  *mockOtel.MockTracer
		mockCtx     context.Context
		mockTraceID string
		mockSpanID  string
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockTracer = mockOtel.NewMockTracer(ctrl)
		mockCtx = context.Background()
		mockTraceID = "58e32b2438efda4d90db7b0ce0db9371"
		mockSpanID = "8fc213b9e9b84a98"
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("API", func() {
		Context("Handler", func() {
			It("should call handler", func() {
				_, mockSpan := mockOtel.MockSpan(mockCtx)
				mockTracer.EXPECT().TracerStart(mockCtx, "handler").Return(mockCtx, mockSpan).Times(1)
				mockTracer.EXPECT().SetIntAttribute(mockCtx, gomock.Any(), gomock.Any()).Times(1)
				mockTracer.EXPECT().TraceID(mockCtx).Return(mockTraceID).Times(1)
				mockTracer.EXPECT().SpanID(mockCtx).Return(mockSpanID).Times(1)

				mockAPI := &TestHandler{
					Tracer: mockTracer,
				}

				mockAPI.Handler(mockCtx)
			})
		})
	})
})
