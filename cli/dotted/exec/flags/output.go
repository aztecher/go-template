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

	JSON bool
	XML  bool
}

func NewOutputFlag() *OutputFlag {
	return &OutputFlag{}
}

func (flag *OutputFlag) Register(ctx context.Context, f *flag.FlagSet) {
	flag.RegisterOnce(func() {
		// general output registration
		f.BoolVar(&flag.JSON, "json", false, "Output JSON format")
		f.BoolVar(&flag.XML, "xml", false, "Output XML format")
	})
}

func (flag *OutputFlag) Process(ctx context.Context) error {
	return flag.ProcessOnce(func() error {
		// general output processing
		return nil
	})
}
