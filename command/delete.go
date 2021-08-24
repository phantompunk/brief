package command

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var deleteUsage = `Removes a specific templates from the saved directory.

Usage: brief delete TEMPLATE

Options:
`

func NewDeleteCommand() *BaseCommand {
	cmd := &BaseCommand{
		flags: flag.NewFlagSet("delete", flag.ExitOnError),
		Execute: func(cmd *BaseCommand, args []string) {
			if len(args) == 0 {
				os.Exit(1)
			}
			file_name := args[0]

			path := filepath.Join("/Users/rodrigomoran/Workspace/brief/template", file_name)
			if _, err := os.Stat(path); err == nil {
				os.Remove(path)
				fmt.Printf("brief: Template '%s' was deleted", file_name)
			}
		},
	}

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, deleteUsage)
	}

	return cmd
}
