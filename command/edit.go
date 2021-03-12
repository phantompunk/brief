package command

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

type EditCommand struct {
	editor   string
	filepath string
	fs       *flag.FlagSet
}

func NewEditCommand() *EditCommand {
	cmd := &EditCommand{
		fs: flag.NewFlagSet("edit", flag.ContinueOnError),
	}

	cmd.fs.StringVar(&cmd.editor, "editor", "vim", "")
	cmd.fs.StringVar(&cmd.filepath, "filepath", "test.md", "")

	cmd.fs.Usage = func() {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("edit usage"))
	}
	return cmd
}

func (c *EditCommand) Init(args []string) error {
	return c.fs.Parse(args)
}

func (c *EditCommand) Called() bool {
	return c.fs.Parsed()
}

func (c *EditCommand) Run() {
	// check for tmp template
	// create if not exists use brief.tmpl as base
	// edit file
	// save
	command := exec.Command(c.editor, c.filepath)
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr
	err := command.Run()
	fmt.Print(err)
}
