package command

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var dest string

func NewAddCommand() *BaseCommand {
	cmd := &BaseCommand{
		flags: flag.NewFlagSet("add", flag.PanicOnError),
		Execute: func(cmd *BaseCommand, args []string) {
			fmt.Printf("Adding %v\n", dest)
			fmt.Println("Running Add()")
			f, err := ioutil.ReadFile(dest)
			if err != nil {
				fmt.Printf("failed to read from %s\n", dest)
				os.Exit(1)
			}

			out, err := os.Create("/Users/rodrigomoran/Workspace/brief/template/test1")
			if err != nil {
				os.Exit(1)
			}
			defer out.Close()

			out.WriteString(string(f))
			fmt.Print("file", string(f))
		},
	}

	cmd.flags.StringVar(&dest, "file", ".brief.tmpl", "destination")

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, "add usage")
	}

	return cmd
}
