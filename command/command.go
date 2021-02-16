package command

import (
	"flag"
	"fmt"
	"os"
	"text/template"
	"time"
)

var createUsage = `Usage: brief create [options...]
Examples:
  # Generate a report for the week containing Feb 2, 2021
	brief create --date 02/17/2021

Options:
  --template	Path to custom template file for weekly report.
  --date	Date used to generate weekly report. Default is current date.
  --output 	Output directory for newly created report. Default is current directory.
`

var versionUsage = `Print the app version and build info for the current context.

Usage: brief version [options...]

Options:
  --short  If true, print just the version number. Default false.
`

var (
	build   = "???"
	version = "???"
)

type Command interface {
	Init()
	Run()
	Called()
}

type VersionCommand struct {
	short bool
	fs    *flag.FlagSet
}

func (c *VersionCommand) Init(args []string) error {
	return c.fs.Parse(args)
}

func (c *VersionCommand) Run() {
	if c.short {
		fmt.Printf("brief version: v%s", version)
	} else {
		fmt.Printf("brief version: v%s, build: %s", version, build)
	}
	os.Exit(0)
}

func (c *VersionCommand) Called() bool {
	return c.fs.Parsed()
}

func NewVersionCommand() *VersionCommand {
	cmd := &VersionCommand{
		fs: flag.NewFlagSet("version", flag.ExitOnError),
	}

	cmd.fs.BoolVar(&cmd.short, "short", false, "")
	cmd.fs.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(versionUsage))
	}
	return cmd
}

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

type weekYear struct {
	Week int
	Year int
	Mon  string
	Tue  string
	Wed  string
	Thu  string
	Fri  string
}

var days = map[int]int{
	0: -1,
	1: 0,
	2: 1,
	3: 2,
	4: 3,
	5: 4,
	6: -2,
}

func getDates(start time.Time) *weekYear {
	year, week := start.ISOWeek()

	firstDayOfWeek := start.AddDate(0, 0, -days[int(start.Weekday())])
	_, m, d := firstDayOfWeek.Date()
	monday := fmt.Sprintf("%d.%d", m, d)

	_, m, d = firstDayOfWeek.AddDate(0, 0, 1).Date()
	tuesday := fmt.Sprintf("%d.%d", m, d)

	_, m, d = firstDayOfWeek.AddDate(0, 0, 2).Date()
	wednesday := fmt.Sprintf("%d.%d", m, d)

	_, m, d = firstDayOfWeek.AddDate(0, 0, 3).Date()
	thursday := fmt.Sprintf("%d.%d", m, d)

	_, m, d = firstDayOfWeek.AddDate(0, 0, 4).Date()
	friday := fmt.Sprintf("%d.%d", m, d)

	return &weekYear{week, year, monday, tuesday, wednesday, thursday, friday}
}
