package command

import (
	"flag"
	"fmt"
	"os"
)

var versionUsage = `Print the app version and build info for the current context.

Usage: brief version [options...]

Options:
  --short  If true, print just the version number. Default false.
`

var (
	build   = "???"
	version = "???"
	short   = false
)

func NewVersionCommand() *BaseCommand {
	cmd := &BaseCommand{
		flags: flag.NewFlagSet("version", flag.ExitOnError),
		Execute: func(cmd *BaseCommand, args []string) {
			if short {
				fmt.Printf("brief version: v%s", version)
			} else {
				fmt.Printf("brief version: v%s, build: %s", version, build)
			}
			os.Exit(0)
		},
	}

	cmd.flags.BoolVar(&short, "short", false, "")

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, versionUsage)
	}

	return cmd
}
