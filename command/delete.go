package command

import (
	"flag"
	"fmt"
	"os"
)

type DeleteCommand struct {
	editor   string
	filepath string
	fs       *flag.FlagSet
}

func NewDeleteCommand() *ListCommand {
	cmd := &ListCommand{
		fs: flag.NewFlagSet("delete", flag.ContinueOnError),
	}

	cmd.fs.StringVar(&cmd.editor, "editor", "vim", "")
	cmd.fs.StringVar(&cmd.filepath, "filepath", "test.md", "")

	cmd.fs.Usage = func() {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("edit usage"))
	}
	return cmd
}

func (c *DeleteCommand) Init(args []string) error {
	return c.fs.Parse(args)
}

func (c *DeleteCommand) Called() bool {
	return c.fs.Parsed()
}

func (c *DeleteCommand) Run() {
	// check for tmp template
	// create if not exists use brief.tmpl as base
	// edit file
	// save
	fmt.Print("Deleting some template")
}
