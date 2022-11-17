package otzap

import (
	"os"
	"strings"
)

// IsInAWSCloud returns true iff the application is running in Amazon Web Services
func IsInAWSCloud() bool {
	return strings.TrimSpace(os.Getenv("AWS_REGION")) != ""

	// Other candidates:
	// - AWS_EXECUTION_ENV
	// - AWS_LAMBDA_FUNCTION_NAME
	// - AWS_LAMBDA_RUNTIME_API
}
