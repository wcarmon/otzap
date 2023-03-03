// MIT License
//
// Copyright (c) 2023 Wilbur Carmon II
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package otzap

// TODO: https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/awsxrayexporter
// TODO: https://docs.aws.amazon.com/xray/latest/devguide/xray-api-segmentdocuments.html
// TODO: https://docs.aws.amazon.com/sdk-for-go/api/service/xray/#XRay.PutTraceSegments
// TODO: https://docs.aws.amazon.com/cli/latest/reference/xray/put-trace-segments.html
// TODO: https://github.com/aws-observability/aws-o11y-recipes/tree/main/sandbox

// TODO: https://github.com/aws-observability/aws-o11y-recipes/blob/main/docs/eks.md
// TODO: https://github.com/aws-observability/aws-otel-collector/blob/main/docs/developers/eks-demo.md

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
