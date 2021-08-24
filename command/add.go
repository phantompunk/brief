package command

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var filePath string
var urlPath string

var addUsage = `Add a template from a file path or URL.

Usage: brief add [OPTIONS] TEMPLATE

Options:
	--file	path to an existing template file
	--url	url for an existing template file
`

func NewAddCommand() *BaseCommand {
	cmd := &BaseCommand{
		flags: flag.NewFlagSet("add", flag.PanicOnError),
		Execute: func(cmd *BaseCommand, args []string) {
			if len(filePath) == 0 {
				fmt.Println("File path required")
				os.Exit(1)
			}

			if _, err := os.Stat(filePath); err != nil {
				fmt.Println("File does not exist")
				os.Exit(1)
			}

			f, err := ioutil.ReadFile(filePath)
			if err != nil {
				fmt.Printf("failed to read from %s\n", filePath)
				os.Exit(1)
			}

			file_name := fmt.Sprint("copy-", filePath)
			if len(args) != 0 {
				file_name = args[0]
			}
			out, err := os.Create("/Users/rodrigomoran/Workspace/brief/template/" + file_name)
			if err != nil {
				os.Exit(1)
			}
			defer out.Close()

			out.WriteString(string(f))
		},
	}

	cmd.flags.StringVar(&filePath, "file", "", "")
	cmd.flags.StringVar(&urlPath, "url", "", "")

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, addUsage)
	}

	return cmd
}
