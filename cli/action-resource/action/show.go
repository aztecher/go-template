package action

import (
	"fmt"
	"flag"
	"github.com/go-template/cli/action-resource/flags"
	"github.com/go-template/cli/action-resource/resource"
)

func init() {
	Register("show", &show{})
}

type show struct {
	*flags.FlagCommon
}

func (s *show) Register(f *flag.FlagSet) {
	s.FlagCommon = flags.NewFlagCommon()
	s.FlagCommon.RegisterOnce(func () {
		// register flag of show action
		fmt.Printf("show flag register\n")
	})
}

func (s *show) Process(resource *resource.Resource) {
	// show resource
}
