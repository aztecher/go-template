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
	baseUsage := BaseUsage()
	baseUsage.Register(fs)

	actions := action.Actions()
	resources := resource.Resources()

	action, cargs, ok := actions.SearchAndConsumeArgument(args)
	if !ok {
		// you can use global level flag usage
		baseUsage.Usage(&actions, &resources, fs)
		fs.Usage()
		return exitHandle.ExitCode
	}

	// action flag assign
	action.Register(fs)

	resource, cargs, ok := resources.SearchAndConsumeArgument(cargs)
	if !ok {
		// you can use action level flag usage
		fmt.Printf("can't match resource name, usage here\n")
		return exitHandle.ExitCode
	}

	// resource flag assign
	resource.Register(fs)

	// register
	if err := fs.Parse(cargs); err != nil {
		fmt.Printf("parse error!!!!!!\n")
		return exitHandle.ExitCode
	}
	fmt.Printf("success parse!!!!!\n")
	fmt.Printf("%+v\n", action)
	fmt.Printf("%+v\n", resource)
	fmt.Printf("cargs: %v\n", cargs)

	action.Process(&resource)
	return exitHandle.ExitCode
}
