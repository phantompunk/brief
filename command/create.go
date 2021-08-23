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

func NewCreateCommand() *BaseCommand {
	cmd := &BaseCommand{
		flags: flag.NewFlagSet("create", flag.PanicOnError),
		Execute: func(cmd *BaseCommand, args []string) {
			file_name := args[0]

			path := filepath.Join("/Users/rodrigomoran/Workspace/brief/template", file_name)
			if _, err := os.Stat(path); err == nil {
				date := time.Now()
				data := getDates(date)

				t, err := template.ParseFiles(path)
				if err != nil {
					os.Exit(1)
				}

				fileName := fmt.Sprintf("Week-%d.md", data.Week)
				f, err := os.Create(fileName)
				if err != nil {
					os.Exit(1)
				}

				err = t.Execute(f, data)
				if err != nil {
					os.Exit(1)
				}
			}

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

// 	cmd.fs.StringVar(&cmd.template, "template", "brief.tmpl", "")
// 	cmd.fs.StringVar(&cmd.date, "date", time.Now().Format("01/02/2006"), "")
// 	cmd.fs.StringVar(&cmd.output, "output", "", "")

// func (c *CreateCommand) Run() {
// 	date, err := time.Parse("01/02/2006", c.date)
// 	if err != nil {
// 		fmt.Println(date)
// 		os.Exit(1)
// 	}

// 	data := getDates(date)

// 	// TODO keep a copy of the template for editing later
// 	t, err := template.ParseFiles(c.template)
// 	if err != nil {
// 		os.Exit(1)
// 	}

// 	fileName := fmt.Sprintf("Week-%d.md", data.Week)
// 	f, err := os.Create(c.output + fileName)
// 	if err != nil {
// 		os.Exit(1)
// 	}
// 	err = t.Execute(f, data)
// 	if err != nil {
// 		os.Exit(1)
// 	}
// }
