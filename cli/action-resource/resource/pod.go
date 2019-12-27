package resource

import (
	"fmt"
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
	pod.FlagCommon = flags.NewFlagCommon()
	pod.FlagCommon.RegisterOnce(func() {
		// register flag of pod resource
		fmt.Printf("pod flag register\n")
	})
}

func (pod *pod) Process() {
}

func (pod *pod) Run() {
}

