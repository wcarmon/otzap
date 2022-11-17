package otzap

import "os"

// IsInGoogleCloud returns true iff the application is running in Google Cloud
func IsInGoogleCloud() bool {
	return os.Getenv("HOME") == "/app"
}
