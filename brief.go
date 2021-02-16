package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/phantompunk/brief/command"
)

var usage = `Usage: brief command [options...]
Options:
  --template	Path to custom template file for weekly report.
  --date	Date used to generate weekly report. Default is current date.
  --output 	Output directory for newly created report. Default is current directory.
Commands:
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

	createCmd := command.NewCreateCommand()
	versionCmd := command.NewVersionCommand()

	switch os.Args[1] {
	case "create":
		createCmd.Init(os.Args[2:])
	case "version":
		versionCmd.Init(os.Args[2:])
	case "edit":
		usageAndExit(fmt.Sprint("brief: 'edit' is not yet implemented.\n"))
	default:
		usageAndExit(fmt.Sprintf("brief: '%s' is not a brief command.\n", os.Args[1]))
	}

	if versionCmd.Called() {
		versionCmd.Run()
	}

	if createCmd.Called() {
		createCmd.Run()
	}
}

func errAndExit(msg string) {
	fmt.Fprintf(os.Stderr, msg)
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
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
