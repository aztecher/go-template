package resource

import (
	"fmt"
	"flag"
	"github.com/go-template/cli/action-resource/flags"
)

func init() {
	Register("empty", &empty{})
}

type empty struct {
	*flags.FlagCommon
}

func (e *empty) Register(f *flag.FlagSet) {
	e.FlagCommon = flags.NewFlagCommon()
	e.FlagCommon.RegisterOnce(func() {
	})
}

func (e *empty) Process() {
}

func (e *empty) ShowProcess() {
	fmt.Printf("empty object!\n")
}

func (e *empty) Run() {
}
