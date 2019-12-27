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
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	// global flag assingn

	actions := action.Actions()
	resources := resource.Resources()

	action, cargs, ok := actions.SearchAndConsumeArgument(args)
	if !ok {
		// you can use global level flag usage
		fmt.Printf("can't match action name, usage here\n")
		return exitHandle.ExitCode
	}

	// action flag assign
	action.Register(fs)

	// resource, ok := resources.Search(args)
	resource, cargs, ok := resources.SearchAndConsumeArgument(cargs)
	if !ok {
		// you can use action level flag usage
		fmt.Printf("can't match resource name, usage here\n")
		return exitHandle.ExitCode
	}

	// resource flag assign
	resource.Register(fs)

	// register
	if err := fs.Parse(args[2:]); err != nil {
		fmt.Printf("parse error!!!!!!\n")
		return exitHandle.ExitCode
	}
	fmt.Printf("success parse!!!!!\n")
	fmt.Printf("%+v\n", action)
	fmt.Printf("%+v\n", resource)
	fmt.Printf("cargs: %v", cargs)
	return exitHandle.ExitCode
}
