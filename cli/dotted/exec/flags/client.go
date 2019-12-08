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

type ClientFlag struct {
	common
}

func NewClientFlag() *ClientFlag {
	return &ClientFlag{}
}

func (flag *ClientFlag) Register(ctx context.Context, f *flag.FlagSet) {
	flag.RegisterOnce(func() {
		// general Registration
	})
}

func (flag *ClientFlag) Process(ctx context.Context) error {
	return flag.ProcessOnce(func() error {
		// general Processing
		return nil
	})
}
