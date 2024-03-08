package main

import (
	"context"

	"fmt"

	"log"

	"net/http"

	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	"go.opentelemetry.io/otel"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"

	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"go.opentelemetry.io/otel/trace"

	"go.uber.org/zap"
)

func initTraceExporter(ctx context.Context) *otlptrace.Exporter {

	traceExporter, err := otlptracegrpc.New(

		ctx,

		otlptracegrpc.WithEndpoint("http://localhost:4317"),
	)

	if err != nil {

		log.Fatalf("failed to create trace exporter: %v", err)

	}

	return traceExporter

}

func initLogExporter(ctx context.Context) *otlptrace.Exporter {

	logExporter, err := otlptracehttp.New(

		ctx,

		otlptracehttp.WithEndpoint("http://localhost:4318/v1/logs"),
	)

	if err != nil {

		log.Fatalf("failed to create log exporter: %v", err)

	}

	return logExporter

}

func initLogger() *zap.Logger {

	logger, err := zap.NewProduction()

	if err != nil {

		log.Fatalf("failed to create logger: %v", err)

	}

	return logger

}
func initTracerProvider(traceExporter *otlptrace.Exporter) *sdktrace.TracerProvider {

	exp, err := stdouttrace.New(stdouttrace.WithPrettyPrint())

	if err != nil {

		log.Println("failed to initialize stdouttrace exporter:", err)

	}

	bsp := sdktrace.NewBatchSpanProcessor(exp)

	tp := sdktrace.NewTracerProvider(

		sdktrace.WithBatcher(traceExporter),

		sdktrace.WithSpanProcessor(bsp),
	)

	return tp

}

func handler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	span := trace.SpanFromContext(ctx)

	defer span.End()

	logger := zap.NewExample().Sugar()

	logger.Infow("Received request",

		"service", "exercise22.02",

		"httpMethod", r.Method,

		"httpURL", r.URL.String(),

		"remoteAddr", r.RemoteAddr,
	)

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Monitoring endpoint invoked!")

}

func main() {

	ctx := context.Background()

	traceExporter := initTraceExporter(ctx)

	defer traceExporter.Shutdown(context.Background())

	logExporter := initLogExporter(ctx)

	defer logExporter.Shutdown(context.Background())

	tp := initTracerProvider(traceExporter)

	otel.SetTracerProvider(tp)

	logger := initLogger()

	defer logger.Sync()
	httpHandler := otelhttp.NewHandler(http.HandlerFunc(handler), "HTTPServer")

	http.Handle("/", httpHandler)

	server := &http.Server{

		Addr: ":8080",

		ReadTimeout: 10 * time.Second,

		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Server listening on port 8080...")

	if err := server.ListenAndServe(); err != nil {

		fmt.Printf("Error starting server: %s\n", err)

	}

}
