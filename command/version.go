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
)

// VersionCommand - Version subcommand
type VersionCommand struct {
	short bool
	fs    *flag.FlagSet
}

// Init - Initializes arguments for Version subcommand
func (c *VersionCommand) Init(args []string) error {
	return c.fs.Parse(args)
}

// Run - Executes core functionallity
func (c *VersionCommand) Run() {
	if c.short {
		fmt.Printf("brief version: v%s", version)
	} else {
		fmt.Printf("brief version: v%s, build: %s", version, build)
	}
	os.Exit(0)
}

// Called - Determines if subcommand has been triggered
func (c *VersionCommand) Called() bool {
	return c.fs.Parsed()
}

// NewVersionCommand - Generates the boilerplate code for a new Command instance
func NewVersionCommand() *VersionCommand {
	cmd := &VersionCommand{
		fs: flag.NewFlagSet("version", flag.ExitOnError),
	}

	cmd.fs.BoolVar(&cmd.short, "short", false, "")
	cmd.fs.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(versionUsage))
	}
	return cmd
}
