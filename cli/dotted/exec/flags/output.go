/*
  Go Project Description

  Date :
  @author Mikiya Michishita
*/

package flags

import (
	"context"
	"flag"
)

type OutputFlag struct {
	common
}

func NewOutputFlag() *OutputFlag {
	return &OutputFlag{}
}

func (flag *OutputFlag) Register(ctx context.Context, f *flag.FlagSet) {
	flag.RegisterOnce(func() {
		// general registration
	})
}

func (flag *OutputFlag) Process(ctx context.Context) error {
	return flag.ProcessOnce(func() error {
		// general processing
		return nil
	})
}
