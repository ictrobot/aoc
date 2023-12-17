//go:build !nogcpercent

package cmd

import (
	"os"
	"runtime/debug"
)

func init() {
	// running all the current solutions in a loop without changing GOGC
	// uses ~32MB, so setting GOGC to 400% for a ~10% performance boost should
	// be safe on almost all systems. This can be overridden by setting the
	// GOGC environment variable, or building with the nogcpercent tag set.
	if _, present := os.LookupEnv("GOGC"); !present {
		debug.SetGCPercent(400)
	}
}
