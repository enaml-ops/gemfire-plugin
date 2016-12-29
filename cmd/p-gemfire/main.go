package main

import (
	pgemfire "github.com/enaml-ops/gemfire-plugin/plugin"
	"github.com/enaml-ops/pluginlib/productv1"
)

// Version is the version of the p-gemfire plugin.
var Version string = "v0.0.0" // overridden at link time

func main() {
	product.Run(&pgemfire.Plugin{
		Version: Version,
	})
}
