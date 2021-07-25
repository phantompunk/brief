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
		fmt.Fprint(os.Stderr, fmt.Sprint(usage))
	}

	flag.Parse()
	if flag.NArg() < 1 {
		usageAndExit("")
	}

	var briefCmd *command.BaseCommand

	switch os.Args[1] {
	case "add":
		briefCmd = command.NewAddCommand()
	case "edit":
		// briefCmd := command.NewEditCommand()
	case "list":
		// briefCmd := command.NewListCommand()
	case "delete":
		// briefCmd := command.NewDeleteCommand()
	case "create":
		// briefCmd := command.NewCreateCommand()
	case "version":
		// briefCmd := command.NewVersionCommand()
	default:
		usageAndExit(fmt.Sprintf("brief: '%s' is not a brief command.\n", os.Args[1]))
	}

	briefCmd.Init(os.Args[2:])
	briefCmd.Run()
}

// func errAndExit(msg string) {
// 	fmt.Fprintf(os.Stderr, msg)
// 	fmt.Fprintf(os.Stderr, "\n")
// 	os.Exit(1)
// }

func usageAndExit(msg string) {
	if msg != "" {
		fmt.Fprint(os.Stderr, msg)
		fmt.Fprintf(os.Stderr, "\n")
	}

	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}
