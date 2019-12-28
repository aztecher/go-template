package cli

import (
	"fmt"
	"os"
	"flag"
	"github.com/go-template/cli/action-resource/flags"
	"github.com/go-template/cli/action-resource/action"
	"github.com/go-template/cli/action-resource/resource"
)

var baseUsage = &BasicUsage{}

type BasicUsage struct {
	option string
	*flags.FlagCommon
}

func (bu *BasicUsage) Register(f *flag.FlagSet) {
	bu.FlagCommon = flags.NewFlagCommon()
	bu.FlagCommon.RegisterOnce(func () {
		f.BoolVar(&bu.FlagCommon.Help, "h", false, "show help")
		// bu.RegBoolVar(f, &bu.FlagCommon.Help, "h", false, "show help")
		// bu.RegBoolVar(f, &bu.FlagCommon.Help, "help", false, "show help")
	})
}

func (bu *BasicUsage) RegBoolVar(f *flag.FlagSet, reg *bool, opt string, def bool, usage string) {
	f.BoolVar(reg, opt, def, usage)
	bu.option = bu.option + usage + "\n"
}

func (bu *BasicUsage) Usage(
	actions *action.ActionMap,
	resources *resource.ResourceMap,
	f *flag.FlagSet,
) {
	f.Parsed()
	f.Usage = func () {
		// custom usage for help
		fmt.Fprintf(os.Stderr, `
Usage of %s
  %s [ACTION] [RESOURCE] OPTIONS...

ACTION:
%s
RESOURCE:
%s
OPTIONS:
`, "commandName", "commandName", actions.List(), resources.List())
		// custom -> bu.RegBoolVar
		f.PrintDefaults() // default output
	}
}

func BaseUsage() *BasicUsage{ 
	return baseUsage
}
