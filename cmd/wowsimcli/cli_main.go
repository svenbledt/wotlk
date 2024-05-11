package main

import (
	"github.com/svenbledt/wotlk/cmd/wowsimcli/cmd"
	"github.com/svenbledt/wotlk/sim"
)

func init() {
	sim.RegisterAll()
}

// Version information.
// This variable is set by the makefile in the release process.
var Version string

func main() {
	cmd.Execute(Version)
}
