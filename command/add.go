package command

import (
	"flag"
	"fmt"
	"os"
)

func NewAddCommand() *BaseCommand {
	cmd := &BaseCommand{
		flags: flag.NewFlagSet("add", flag.ContinueOnError),
		Execute: func(cmd *BaseCommand, args []string) {
			fmt.Print("Running Add()")
		},
	}

	cmd.flags.Usage = func() {
		fmt.Fprint(os.Stderr, "add usage")
	}

	return cmd
}
