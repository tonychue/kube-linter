// +build tools

package toolimports

// This file declares dependencies on tool for `go mod` purposes.
// See https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
// for an explanation of the approach.

import (
	_ "github.com/gobuffalo/packr/packr"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "honnef.co/go/tools/cmd/staticcheck"
)