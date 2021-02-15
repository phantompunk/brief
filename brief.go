package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"
	"time"
)

// type CreateCommand struct {
// 	date     string
// 	template string
// 	fs       *flag.FlagSet
// }

// func NewCreateCommand() *CreateCommand {
// 	createCmd := &CreateCommand{
// 		fs: flag.NewFlagSet("create", flag.ContinueOnError),
// 	}

// 	createCmd.fs.StringVar(&createCmd.template, "template", "brief.tmpl", "Path to custom template")
// 	createCmd.fs.StringVar(&createCmd.date, "date", time.Now().Format("02-14-2021"), "Given date")
// 	createCmd.fs.Parse(os.Args[1:])
// 	fmt.Println(createCmd.date)
// 	fmt.Println(createCmd.template)
// 	return createCmd
// }

// func main() {
// 	NewCreateCommand()
// }

type WeekYear struct {
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

func getDates(start time.Time) *WeekYear {
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

	return &WeekYear{week, year, monday, tuesday, wednesday, thursday, friday}
}

var usage = `Usage: brief [options...] command <url>
Options:
  --template	Path to custom template file for weekly report.
  --date	Date used to generate weekly report. Default is current date.
  --output 	Output directory for newly created report. Default is current directory.
Subcommands:
  create	Generate a weekly report
  edit		Edit weekly report template
  version	Print version info
`

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage))
	}

	flag.Parse()
	if flag.NArg() < 1 {
		usageAndExit("")
	}

	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	createTempl := createCmd.String("template", "brief.tmpl", "Generate weekly report using custom template")
	createDate := createCmd.String("date", "2012/01/24", "Generate weekly report for a given date")
	createOut := createCmd.String("output", "", "Path to place weekly report")

	if len(os.Args) < 2 {
		fmt.Println("Usage: brief create [OPTIONS]")
		createCmd.PrintDefaults()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create":
		createCmd.Parse(os.Args[2:])
		templatePath := *createTempl
		outPath := *createOut
		date, err := time.Parse("2006/01/02", *createDate)
		if err != nil {
			fmt.Println(date)
			os.Exit(1)
		}

		fmt.Println("SubCommand: Create")
		fmt.Println("Template:", templatePath)
		fmt.Println("Date:", date)
		fmt.Println("Args", createCmd.Args())

		data := getDates(date)

		t, err := template.ParseFiles(templatePath)
		if err != nil {
			os.Exit(1)
		}

		fileName := fmt.Sprintf("Week-%d.md", data.Week)
		f, err := os.Create(outPath + fileName)
		if err != nil {
			os.Exit(1)
		}
		err = t.Execute(f, data)
	default:
		usageAndExit(fmt.Sprintf("brief: '%s' is not a brief command.\n", os.Args[1]))
	}
}

func usageAndExit(msg string) {
	if msg != "" {
		fmt.Fprintf(os.Stderr, msg)
		fmt.Fprintf(os.Stderr, "\n")
	}

	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}