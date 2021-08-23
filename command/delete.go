package command

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func NewDeleteCommand() *BaseCommand {
	cmd := &BaseCommand{
		flags: flag.NewFlagSet("delete", flag.PanicOnError),
		Execute: func(cmd *BaseCommand, args []string) {
			file_name := args[0]
			fmt.Printf("Deleting file %v", file_name)

			path := filepath.Join("/Users/rodrigomoran/Workspace/brief/template", file_name)
			if _, err := os.Stat(path); err == nil {
				os.Remove(path)
				fmt.Printf("brief: removed %s", file_name)
			}
			// file, err := os.Open("/Users/rodrigomoran/Workspace/brief/template")
			// if err != nil {
			// 	os.Exit(1)
			// }
			// defer file.Close()

			// filelist, _ := file.Readdir(0)

			// fmt.Printf("Name\t\tSize\t\tModified\n")
			// for _, files := range filelist {
			// 	fmt.Printf("\n%-15s %-7v %v", files.Name(), files.Size(), files.ModTime())
			// }
		},
	}

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, "add usage")
	}

	return cmd
}
