package cli

import (
	"fmt"
	"flag"
	"github.com/go-template/cli/action-resource/action"
	"github.com/go-template/cli/action-resource/resource"
	"github.com/go-template/cli/action-resource/exitcode"
)

type Command interface {
	Run ()
}

func Run(args[]string) int{
	exitHandle := exitcode.NewHandle().DefaultErr()

	actions := action.Actions()
	resources := resource.Resources()

	action, ok := actions.Search(args)
	if !ok {
		fmt.Printf("can't match action name, usage here\n")
		return exitHandle.ExitCode
	}

	resource, ok := resources.Search(args)
	if !ok {
		fmt.Printf("can't match resource name, usage here\n")
		return exitHandle.ExitCode
	}

	fmt.Printf("%+v\n", action)
	fmt.Printf("%+v\n", resource)

	// subcommand
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	// register
	if err := fs.Parse(args); err != nil {
		fmt.Printf("parse error!!!!!!\n")
		return exitHandle.ExitCode
	}
	fmt.Printf("success parse!!!!!\n")
	return exitHandle.ExitCode
}
