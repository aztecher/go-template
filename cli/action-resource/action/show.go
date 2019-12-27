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

	JSON bool
}

func (s *show) Register(f *flag.FlagSet) {
	s.FlagCommon = flags.NewFlagCommon()
	s.FlagCommon.RegisterOnce(func () {
		// register flag of show action
		fmt.Printf("show flag register\n")
	})

	f.BoolVar(&s.JSON, "j", false, "Output JSON format")
	f.BoolVar(&s.JSON, "json", false, "Output JSON format")
}

func (s *show) Process(resource *resource.Resource) {
	// show resource
}
