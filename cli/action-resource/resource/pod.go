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
}

func (pod *pod) Register(f *flag.FlagSet) {
	pod.FlagCommon = flags.NewFlagCommon()
	pod.FlagCommon.RegisterOnce(func() {
	})
}

func (pod *pod) Process() {
	// ここでresourceの初期化とか？
	// Runで必要なものの準備とか
}

func (pod *pod) ShowProcess() {
	fmt.Printf("pod process (calling from XXX)\n")
}

func (pod *pod) Run() {
}

