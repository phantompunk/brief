package command

import (
	"flag"
	"fmt"
	"os"
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

func NewCreateCommand() *CreateCommand {
	cmd := &CreateCommand{
		fs: flag.NewFlagSet("create", flag.ContinueOnError),
	}

	cmd.fs.StringVar(&cmd.template, "template", "brief.tmpl", "")
	cmd.fs.StringVar(&cmd.date, "date", time.Now().Format("01/02/2006"), "")
	cmd.fs.StringVar(&cmd.output, "output", "", "")

	cmd.fs.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(createUsage))
	}
	return cmd
}

func (c *CreateCommand) Init(args []string) error {
	return c.fs.Parse(args)
}

func (c *CreateCommand) Called() bool {
	return c.fs.Parsed()
}

func (c *CreateCommand) Run() {
	date, err := time.Parse("01/02/2006", c.date)
	if err != nil {
		fmt.Println(date)
		os.Exit(1)
	}

	data := getDates(date)

	// TODO keep a copy of the template for editing later
	t, err := template.ParseFiles(c.template)
	if err != nil {
		os.Exit(1)
	}

	fileName := fmt.Sprintf("Week-%d.md", data.Week)
	f, err := os.Create(c.output + fileName)
	if err != nil {
		os.Exit(1)
	}
	err = t.Execute(f, data)
}
