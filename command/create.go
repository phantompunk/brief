package command

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	"time"
)

type Command interface {
	Init()
	Run()
	Called()
}

var createUsage = `Usage: brief create [options...]
Examples:
  # Generate a report for the week containing Feb 2, 2021
	brief create --date 02/17/2021

Options:
  --template	Path to custom template file for weekly report.
  --date	Date used to generate weekly report. Default is current date.
  --output 	Output directory for newly created report. Default is current directory.
`

type CreateCommand struct {
	date     string
	template string
	output   string
	fs       *flag.FlagSet
}

var dest string

func NewCreateCommand() *BaseCommand {
	cmd := &BaseCommand{
		flags: flag.NewFlagSet("create", flag.PanicOnError),
		Execute: func(cmd *BaseCommand, args []string) {
			file_name := args[0]

			path := filepath.Join("/Users/rodrigomoran/Workspace/brief/template/", file_name)
			if _, err := os.Stat(path); err == nil {
				fmt.Print("file exists")
				date := time.Now()
				data := getDates(date)

				t, err := template.ParseFiles(path)
				if err != nil {
					fmt.Print("Failed to parse")
					os.Exit(1)
				}

				fileName := fmt.Sprintf("Week-%d.md", data.Week)
				f, err := os.Create(fileName)
				if err != nil {
					fmt.Print("Failed to create")
					os.Exit(1)
				}

				err = t.Execute(f, data)
				if err != nil {
					fmt.Print("Failed to execute")
					os.Exit(1)
				}
			}
			fmt.Print("What happened?")

			// f, err := ioutil.ReadFile(dest)
			// if err != nil {
			// 	fmt.Printf("failed to read from %s\n", dest)
			// 	os.Exit(1)
			// }

			// file_name := fmt.Sprint("copy-", dest)
			// if len(args) != 0 {
			// 	file_name = args[0]
			// }
			// out, err := os.Create("/Users/rodrigomoran/Workspace/brief/template/" + file_name)
			// if err != nil {
			// 	os.Exit(1)
			// }
			// defer out.Close()

			// out.WriteString(string(f))
		},
	}

	cmd.flags.StringVar(&dest, "dest", ".", "destination")

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, "add usage")
	}

	return cmd
}
