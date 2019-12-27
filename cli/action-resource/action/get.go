package action

import (
	"fmt"
	"flag"
	"github.com/go-template/cli/action-resource/flags"
	"github.com/go-template/cli/action-resource/resource"
)

func init() {
	Register("get", &get{})
}

type get struct {
	*flags.FlagCommon
}

func (g *get) Register(f *flag.FlagSet) {
	g.FlagCommon = flags.NewFlagCommon()
	g.FlagCommon.RegisterOnce(func () {
		// register flag of get action
		fmt.Printf("get flag register\n")
	})
}

func (g *get) Process(resource *resource.Resource) {
	// getting resources
}
