package resource

import (
	"flag"
	"github.com/go-template/cli/action-resource/flags"
)

func init() {
	Register("pod", &pod{})
}

type pod struct {
	*flags.FlagCommon
	name string
}

func (pod *pod) Register(f *flag.FlagSet) {
	pod.FlagCommon.RegisterOnce(func() {
		// register flag of pod resource
	})
}

func (pod *pod) Process() {
}

func (pod *pod) Run() {
}

