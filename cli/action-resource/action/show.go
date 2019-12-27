package action

import (
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
	s.FlagCommon.RegisterOnce(func () {
		// register flag of show action
	})
}

func (s *show) Process(resource *resource.Resource) {
	// show resource
}
