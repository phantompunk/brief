package command

import (
	"flag"
	"fmt"
	"os"
)

var listUsage = `List all currently avaible templates.

Usage: brief list
Options:
`

func NewListCommand() *BaseCommand {
	cmd := &BaseCommand{
		flags: flag.NewFlagSet("list", flag.ExitOnError),
		Execute: func(cmd *BaseCommand, args []string) {
			file, err := os.Open("/Users/rodrigomoran/Workspace/brief/template")
			if err != nil {
				os.Exit(1)
			}
			defer file.Close()

			filelist, _ := file.Readdir(0)

			fmt.Printf("NAME\t\tSIZE\t\tMODIFIED")
			for _, files := range filelist {
				fmt.Printf("\n%-15s %-15v %v", files.Name(), files.Size(), files.ModTime().Format("2006-01-02 15:04:05"))
			}
		},
	}

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, listUsage)
	}

	return cmd
}

// type ListCommand struct {
// 	editor   string
// 	filepath string
// 	fs       *flag.FlagSet
// }

// func NewListCommand() *ListCommand {
// 	cmd := &ListCommand{
// 		fs: flag.NewFlagSet("list", flag.ContinueOnError),
// 	}

// 	cmd.fs.StringVar(&cmd.editor, "editor", "vim", "")
// 	cmd.fs.StringVar(&cmd.filepath, "filepath", "test.md", "")

// 	cmd.fs.Usage = func() {
// 		fmt.Fprint(os.Stderr, "edit usage")
// 	}
// 	return cmd
// }

// func (c *ListCommand) Init(args []string) error {
// 	return c.fs.Parse(args)
// }

// func (c *ListCommand) Called() bool {
// 	return c.fs.Parsed()
// }

// func (c *ListCommand) Run() {
// 	// check for tmp template
// 	// create if not exists use brief.tmpl as base
// 	// edit file
// 	// save
// 	fmt.Print("List some templates")
// }
