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

import (
	"go.uber.org/zap"
	"os"
	"strings"
)

// PrintEnvVars prints local environment variables using zap API
func PrintEnvVars() {
	fields := make([]zap.Field, 0, len(os.Environ()))
	fields = append(fields, zap.Int("count", len(os.Environ())))

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fields = append(fields, zap.String(pair[0], pair[1]))
	}

	zap.L().Info("env vars", fields...)
}
