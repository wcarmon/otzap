package otzap

// TODO: https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/awsxrayexporter
// TODO: https://docs.aws.amazon.com/xray/latest/devguide/xray-api-segmentdocuments.html
// TODO: https://docs.aws.amazon.com/sdk-for-go/api/service/xray/#XRay.PutTraceSegments
// TODO: https://docs.aws.amazon.com/cli/latest/reference/xray/put-trace-segments.html

// TODO: AWS X-Ray requires their ID generator

/*


traceExporter, err := otlptracegrpc.New(
	ctx,
	otlptracegrpc.WithInsecure(),
	otlptracegrpc.WithEndpoint(endpoint),
	otlptracegrpc.WithDialOption(grpc.WithBlock())

// See https://pkg.go.dev/go.opentelemetry.io/contrib/propagators/aws/xray#IDGenerator
idGen := xray.NewIDGenerator()

tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(traceExporter),
		sdktrace.WithIDGenerator(idg),
		...

func getXrayTraceID(span trace.Span) string {
	xrayTraceID := span.SpanContext().TraceID().String()
	result := fmt.Sprintf("1-%s-%s", xrayTraceID[0:8], xrayTraceID[8:])
	return result

*/
