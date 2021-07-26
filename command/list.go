package command

import (
	"flag"
	"fmt"
	"os"
)

type ListCommand struct {
	editor   string
	filepath string
	fs       *flag.FlagSet
}

func NewListCommand() *ListCommand {
	cmd := &ListCommand{
		fs: flag.NewFlagSet("list", flag.ContinueOnError),
	}

	cmd.fs.StringVar(&cmd.editor, "editor", "vim", "")
	cmd.fs.StringVar(&cmd.filepath, "filepath", "test.md", "")

	cmd.fs.Usage = func() {
		fmt.Fprint(os.Stderr, "edit usage")
	}
	return cmd
}

func (c *ListCommand) Init(args []string) error {
	return c.fs.Parse(args)
}

func (c *ListCommand) Called() bool {
	return c.fs.Parsed()
}

func (c *ListCommand) Run() {
	// check for tmp template
	// create if not exists use brief.tmpl as base
	// edit file
	// save
	fmt.Print("List some templates")
}
