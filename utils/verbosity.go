package utils

import (
	"github.com/spf13/pflag"
)

// Verbose flag. Use pflag.Parse() in main() to populate.
var Verbose bool

func init() {
	pflag.BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}
