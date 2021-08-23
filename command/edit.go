package command

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// type EditCommand struct {
// 	editor   string
// 	filepath string
// 	fs       *flag.FlagSet
// }

func NewEditCommand() *BaseCommand {
	cmd := &BaseCommand{
		flags: flag.NewFlagSet("edit", flag.ContinueOnError),
		Execute: func(cmd *BaseCommand, args []string) {
			file_name := args[0]
			file_path := filepath.Join("/Users/rodrigomoran/Workspace/brief/template", file_name)

			if _, err := os.Stat(file_path); err == nil {
				command := exec.Command("vim", file_path)
				command.Stdout = os.Stdout
				command.Stdin = os.Stdin
				command.Stderr = os.Stderr
				err := command.Run()
				fmt.Print(err)
			}
			fmt.Printf("Running Edit() on %s", file_name)
		},
	}

	cmd.flags.Usage = func() {
		fmt.Fprint(os.Stderr, "add usage")
	}

	return cmd
}
