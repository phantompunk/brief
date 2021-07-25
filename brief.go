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

	addCmd := command.NewAddCommand()
	editCmd := command.NewEditCommand()
	listCmd := command.NewListCommand()
	deleteCmd := command.NewDeleteCommand()
	createCmd := command.NewCreateCommand()
	versionCmd := command.NewVersionCommand()

	switch os.Args[1] {
	case "add":
		// usageAndExit(fmt.Sprint("brief: 'edit' is not yet implemented.\n"))
		addCmd.Init(os.Args[2:])
		addCmd.Run()
	case "edit":
		// usageAndExit(fmt.Sprint("brief: 'edit' is not yet implemented.\n"))
		editCmd.Init(os.Args[2:])
		editCmd.Run()
	case "list":
		// usageAndExit(fmt.Sprint("brief: 'edit' is not yet implemented.\n"))
		listCmd.Init(os.Args[2:])
		listCmd.Run()
	case "delete":
		// usageAndExit(fmt.Sprint("brief: 'edit' is not yet implemented.\n"))
		deleteCmd.Init(os.Args[2:])
		deleteCmd.Run()
	case "create":
		createCmd.Init(os.Args[2:])
		createCmd.Run()
	case "version":
		versionCmd.Init(os.Args[2:])
		versionCmd.Run()
	default:
		usageAndExit(fmt.Sprintf("brief: '%s' is not a brief command.\n", os.Args[1]))
	}
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
