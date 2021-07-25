package command

import (
	"flag"
	"fmt"
	"os"
)

type AddCommand struct {
	editor   string
	filepath string
	fs       *flag.FlagSet
}

func NewAddCommand() *AddCommand {
	cmd := &AddCommand{
		fs: flag.NewFlagSet("edit", flag.ContinueOnError),
	}

	cmd.fs.StringVar(&cmd.editor, "editor", "vim", "")
	cmd.fs.StringVar(&cmd.filepath, "filepath", "test.md", "")

	cmd.fs.Usage = func() {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("edit usage"))
	}
	return cmd
}

func (c *AddCommand) Init(args []string) error {
	return c.fs.Parse(args)
}

func (c *AddCommand) Called() bool {
	return c.fs.Parsed()
}

func (c *AddCommand) Run() {
	// check for tmp template
	// create if not exists use brief.tmpl as base
	// edit file
	// save
	fmt.Print("Adding some template")
}
