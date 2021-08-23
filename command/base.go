package command

import (
	"flag"
	"fmt"
)

type BaseCommand struct {
	flags   *flag.FlagSet
	Execute func(cmd *BaseCommand, args []string)
}

func (c *BaseCommand) Init(args []string) error {
	fmt.Printf("Init Args: %v\n", args)
	return c.flags.Parse(args)
}

func (c *BaseCommand) Called() bool {
	fmt.Println("Parsing Command")
	return c.flags.Parsed()
}

func (c *BaseCommand) Run() {
	fmt.Println("Running Command")
	c.Execute(c, c.flags.Args())
}
