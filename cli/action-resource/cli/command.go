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

// example of bellow commands
// $ command action resource [options]
// action resourceが与えられると、Processingするようにしている
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

// 実はresourceがないようなコマンドの場合にも対応できる.
// empty resourceを作成し、それにresourceがない場合の処理を記述すれば良い
// あとは, empty resourceを利用するように, resourceのsearch後を調整する
func RunWithEmpty(args[]string) int{
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
		// ここの変更のみ
		empty := []string{"empty"}
		resource, _, _ = resources.SearchAndConsumeArgument(empty)
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
