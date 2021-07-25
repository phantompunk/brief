package command

import "flag"

type BaseCommand struct {
	flags   *flag.FlagSet
	Execute func(cmd *BaseCommand, args []string)
}

func (c *BaseCommand) Init(args []string) error {
	return c.flags.Parse(args)
}

func (c *BaseCommand) Called() bool {
	return c.flags.Parsed()
}

func (c *BaseCommand) Run() {
	c.Execute(c, c.flags.Args())
}
