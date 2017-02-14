package main

import (
	"github.com/docopt/docopt-go"
)

const version = "0.0.0"

const usage = `
Flare is a lightweight self discoving monitoring solution.

Usage: flare [-h --help]
       flare [--discover|--version]

Options:
  -h --help    Show this message
  --version    Show the version number and exit
  --discover   Run only in discovery mode output neighbors found.
`

// Args are arguments from the command line.
type Args struct {
	Discover bool
}

// ParseArgs parses the command line arguments.
func ParseArgs() *Args {
	arguments, _ := docopt.Parse(usage, nil, true, version, true, true)

	args := Args{
		Discover: arguments["--discover"].(bool),
	}

	return &args
}
